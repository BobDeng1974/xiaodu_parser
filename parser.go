package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	devProto "xiaodu_parser/dev_proto"
	pb "xiaodu_parser/grpc_proto"
)
const (
	Category_CMD    uint32 = 0
	Category_CONFIG uint32 = 1
)
const (
	TYPE_CMD = "cmd"
	CMD_LAMP = "lamp"
	CMD_ON = "on"
	CMD_OFF = "off"
	TYPE_CONFIG = "config"
	CONFIG_VOLUME = "VOLUME"
)

type server struct{}

func (this *server) Marshal(ctx context.Context, in *pb.DownReq) (out *pb.DownRsp, err error) {
	log.Println("[Marshal]", in.Name)
	if ctx.Err() == context.Canceled {
		log.Println("timeout")
		return nil,errors.New ("[Marshal]timeout")
	}
	//应用层下发的指令可能有多种组合
	downMsg := &devProto.Payload{}
	kind := in.Kind
	filed := in.Field
	val := in.Val
	if kind == TYPE_CMD{
		downMsg.Kind = uint32(devProto.Category_CMD)
		switch filed {
		case CMD_LAMP:
			// 下发给灯的值，是保留几位小数，要不要扩大10倍数。要不要把下发的值范围由100~1 量化成中控的范围10~1等逻辑处理等等
			// 应该交给解析器完成。既不要交给上层应用，因为他们没必要了解终端逻辑和数据类型。也不要交给终端完成这些，因为能在解
			// 析器完成，就尽量不要在终端。因为涉及日后更新，bug修复等，终端改改起来成本都很高。
			downMsg.Key = uint32(devProto.Device_LAMP)
			if mustString(val) == CMD_ON{
				downMsg.Val = []byte{uint8(devProto.Operation_ON)}
			}else {
				downMsg.Val=[]byte{uint8(devProto.Operation_OFF)}
			}
			payload , err :=CreateMarshalDown(downMsg)
			if err != nil {
				log.Println("[Marshal]CreateMarshalDown error",err)
				err = errors.Wrap(err,"CreateMarshalDown error")
				return nil ,err
			}
			out = &pb.DownRsp{}
			out.ID = in.ID
			out.Name = in.Name
			out.Payload = payload

			return  out,nil
		default:
			log.Println("[Marshal]filed=", filed)
			return nil,errors.New ("[Marshal]cmd filed error")
		}

	}else if kind == TYPE_CONFIG {
			//TODO:配置参数的修改
		log.Println("not support config", kind)

	}else{

		log.Println("kind error.", kind)
		return nil,errors.New ("[Marshal]kind error")
	}
	log.Println("Unknown error ")
	err = errors.New("Unknown error")
	return nil,err
}

// Name 字段表示不同设备定制厂商名称。同样一款产品，定制厂商要求配置可能不一样的。解析结果也不一样
// 对于不同版本解析方式的不同，通过payload的首字段来表示版本号
func (this *server) UnMarshal(ctx context.Context, in *pb.UpReq) (out *pb.UpRsp, err error) {
	log.Println("[Marshal]", in.Name)
	if ctx.Err() == context.Canceled {
		return nil,errors.New ("[Marshal]timeout")
	}
	payDecrypt,err := decrypt(in.Payload)
	if err != nil {
		log.Println("[controlHandle]Decrypt error",err)
		return nil,err
	}
	upPayload := &devProto.Payload{}
	err = proto.Unmarshal(payDecrypt, upPayload)
	if err != nil {
		log.Println("[controlHandle] proto Unmarshal error",err)
		return  nil ,err
	}
	out = &pb.UpRsp{}
	out.ID = in.ID
	out.Name = in.Name
	out.Kind = TYPE_CMD
	out.Field = CMD_LAMP

	if 	upPayload.Val[0] == byte(devProto.Operation_ON){
		out.Val = "on"
	}else {
		out.Val = "off"
	}
	return out, nil
}

func CreateNewParser() error{
	ln, err := net.Listen("tcp", ":"+ strconv.Itoa(PORT))
	if err != nil {
		log.Println("网络错误: ", err)
		ConsulDeRegister()

		return errors.Wrap(err,"network error")
	}

	//创建grpc的服务
	srv := grpc.NewServer()
	//注册服务
	pb.RegisterParserServer(srv, &server{})
	//等待网络连接
	err = srv.Serve(ln)
	if err != nil {
		ConsulDeRegister()
		log.Println("网络错误: ", err)
		return errors.Wrap(err,"connect error")
	}
	// TODO:如果服务建立出现问题，应该将注册的服务主动停掉
	return nil
}

func mustString(data interface{}) string {
	if data != nil && fmt.Sprintf("%T", data) == "string" {
		return data.(string)
	}
	return ""
}