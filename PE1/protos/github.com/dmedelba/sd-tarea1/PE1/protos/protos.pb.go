// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Solicitamos un pedido
type SolicitudPedidoPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//id-paquete tipo nombre valor origen destino
	IdPaquete string `protobuf:"bytes,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	Tipo      string `protobuf:"bytes,2,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Nombre    string `protobuf:"bytes,3,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Valor     string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Origen    string `protobuf:"bytes,5,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino   string `protobuf:"bytes,6,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *SolicitudPedidoPyme) Reset() {
	*x = SolicitudPedidoPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudPedidoPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudPedidoPyme) ProtoMessage() {}

func (x *SolicitudPedidoPyme) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudPedidoPyme.ProtoReflect.Descriptor instead.
func (*SolicitudPedidoPyme) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{0}
}

func (x *SolicitudPedidoPyme) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *SolicitudPedidoPyme) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *SolicitudPedidoPyme) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *SolicitudPedidoPyme) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *SolicitudPedidoPyme) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *SolicitudPedidoPyme) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type SolicitudPedidoRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//id-paquete tipo nombre valor origen destino
	IdPaquete string `protobuf:"bytes,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	Nombre    string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Valor     string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Origen    string `protobuf:"bytes,4,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino   string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *SolicitudPedidoRetail) Reset() {
	*x = SolicitudPedidoRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudPedidoRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudPedidoRetail) ProtoMessage() {}

func (x *SolicitudPedidoRetail) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudPedidoRetail.ProtoReflect.Descriptor instead.
func (*SolicitudPedidoRetail) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{1}
}

func (x *SolicitudPedidoRetail) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *SolicitudPedidoRetail) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *SolicitudPedidoRetail) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *SolicitudPedidoRetail) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *SolicitudPedidoRetail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

//respuesta al solicitar un pedido tanto pyme/retail
type RespuestaPedido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodigoSeguimiento string `protobuf:"bytes,1,opt,name=codigoSeguimiento,proto3" json:"codigoSeguimiento,omitempty"`
}

func (x *RespuestaPedido) Reset() {
	*x = RespuestaPedido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaPedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaPedido) ProtoMessage() {}

func (x *RespuestaPedido) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaPedido.ProtoReflect.Descriptor instead.
func (*RespuestaPedido) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{2}
}

func (x *RespuestaPedido) GetCodigoSeguimiento() string {
	if x != nil {
		return x.CodigoSeguimiento
	}
	return ""
}

//Solicitamos un seguimiento
type SolicitudSeguimiento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//id-paquete tipo nombre valor origen destino
	CodigoSeguimiento string `protobuf:"bytes,1,opt,name=codigoSeguimiento,proto3" json:"codigoSeguimiento,omitempty"`
}

func (x *SolicitudSeguimiento) Reset() {
	*x = SolicitudSeguimiento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudSeguimiento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudSeguimiento) ProtoMessage() {}

func (x *SolicitudSeguimiento) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudSeguimiento.ProtoReflect.Descriptor instead.
func (*SolicitudSeguimiento) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{3}
}

func (x *SolicitudSeguimiento) GetCodigoSeguimiento() string {
	if x != nil {
		return x.CodigoSeguimiento
	}
	return ""
}

//Respuesta del estado
type RespuestaSeguimiento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//id-paquete tipo nombre valor origen destino
	EstadoPedido string `protobuf:"bytes,1,opt,name=estadoPedido,proto3" json:"estadoPedido,omitempty"`
}

func (x *RespuestaSeguimiento) Reset() {
	*x = RespuestaSeguimiento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaSeguimiento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaSeguimiento) ProtoMessage() {}

func (x *RespuestaSeguimiento) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaSeguimiento.ProtoReflect.Descriptor instead.
func (*RespuestaSeguimiento) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{4}
}

func (x *RespuestaSeguimiento) GetEstadoPedido() string {
	if x != nil {
		return x.EstadoPedido
	}
	return ""
}

type SolicitudPaquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo      string `protobuf:"bytes,1,opt,name=tipo,proto3" json:"tipo,omitempty"`
	IdPaquete string `protobuf:"bytes,2,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
}

func (x *SolicitudPaquete) Reset() {
	*x = SolicitudPaquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudPaquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudPaquete) ProtoMessage() {}

func (x *SolicitudPaquete) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudPaquete.ProtoReflect.Descriptor instead.
func (*SolicitudPaquete) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{5}
}

func (x *SolicitudPaquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *SolicitudPaquete) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

//respuesta solicitud paquete
type RespuestaPaquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete         string `protobuf:"bytes,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	CodigoSeguimiento int32  `protobuf:"varint,2,opt,name=codigoSeguimiento,proto3" json:"codigoSeguimiento,omitempty"`
	Tipo              string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor             int32  `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Origen            string `protobuf:"bytes,5,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino           string `protobuf:"bytes,6,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *RespuestaPaquete) Reset() {
	*x = RespuestaPaquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaPaquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaPaquete) ProtoMessage() {}

func (x *RespuestaPaquete) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaPaquete.ProtoReflect.Descriptor instead.
func (*RespuestaPaquete) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{6}
}

func (x *RespuestaPaquete) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *RespuestaPaquete) GetCodigoSeguimiento() int32 {
	if x != nil {
		return x.CodigoSeguimiento
	}
	return 0
}

func (x *RespuestaPaquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *RespuestaPaquete) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *RespuestaPaquete) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *RespuestaPaquete) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

//estado del pedido
type SolicitudEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete string `protobuf:"bytes,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	Intentos  string `protobuf:"bytes,2,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado    string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
	Fecha     string `protobuf:"bytes,4,opt,name=fecha,proto3" json:"fecha,omitempty"`
}

func (x *SolicitudEstado) Reset() {
	*x = SolicitudEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudEstado) ProtoMessage() {}

func (x *SolicitudEstado) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudEstado.ProtoReflect.Descriptor instead.
func (*SolicitudEstado) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{7}
}

func (x *SolicitudEstado) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *SolicitudEstado) GetIntentos() string {
	if x != nil {
		return x.Intentos
	}
	return ""
}

func (x *SolicitudEstado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *SolicitudEstado) GetFecha() string {
	if x != nil {
		return x.Fecha
	}
	return ""
}

type RespuestaEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmacion string `protobuf:"bytes,1,opt,name=confirmacion,proto3" json:"confirmacion,omitempty"`
}

func (x *RespuestaEstado) Reset() {
	*x = RespuestaEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaEstado) ProtoMessage() {}

func (x *RespuestaEstado) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaEstado.ProtoReflect.Descriptor instead.
func (*RespuestaEstado) Descriptor() ([]byte, []int) {
	return file_protos_proto_rawDescGZIP(), []int{8}
}

func (x *RespuestaEstado) GetConfirmacion() string {
	if x != nil {
		return x.Confirmacion
	}
	return ""
}

var File_protos_proto protoreflect.FileDescriptor

var file_protos_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x73, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x75, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x22, 0x95, 0x01, 0x0a, 0x15, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x50, 0x65,
	0x64, 0x69, 0x64, 0x6f, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64,
	0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x3f, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x63,
	0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65,
	0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x14, 0x73, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74,
	0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69,
	0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f,
	0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x73,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x70, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75,
	0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65,
	0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x79,
	0x0a, 0x0f, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65, 0x63, 0x68, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x65, 0x63, 0x68, 0x61, 0x22, 0x35, 0x0a, 0x0f, 0x72, 0x65, 0x73,
	0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x32, 0x93, 0x03, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x4d, 0x0a, 0x13, 0x73,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x50, 0x79,
	0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x75, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x15, 0x73, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x58, 0x0a,
	0x18, 0x6f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65,
	0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d,
	0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x61, 0x72, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x50, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0d, 0x6f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x75, 0x64, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x65, 0x64, 0x65, 0x6c, 0x62, 0x61, 0x2f, 0x73, 0x64,
	0x2d, 0x74, 0x61, 0x72, 0x65, 0x61, 0x31, 0x2f, 0x50, 0x45, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_proto_rawDescOnce sync.Once
	file_protos_proto_rawDescData = file_protos_proto_rawDesc
)

func file_protos_proto_rawDescGZIP() []byte {
	file_protos_proto_rawDescOnce.Do(func() {
		file_protos_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_proto_rawDescData)
	})
	return file_protos_proto_rawDescData
}

var file_protos_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_proto_goTypes = []interface{}{
	(*SolicitudPedidoPyme)(nil),   // 0: protos.solicitudPedidoPyme
	(*SolicitudPedidoRetail)(nil), // 1: protos.solicitudPedidoRetail
	(*RespuestaPedido)(nil),       // 2: protos.respuestaPedido
	(*SolicitudSeguimiento)(nil),  // 3: protos.solicitudSeguimiento
	(*RespuestaSeguimiento)(nil),  // 4: protos.respuestaSeguimiento
	(*SolicitudPaquete)(nil),      // 5: protos.solicitudPaquete
	(*RespuestaPaquete)(nil),      // 6: protos.respuestaPaquete
	(*SolicitudEstado)(nil),       // 7: protos.solicitudEstado
	(*RespuestaEstado)(nil),       // 8: protos.respuestaEstado
}
var file_protos_proto_depIdxs = []int32{
	0, // 0: protos.protos.solicitarPedidoPyme:input_type -> protos.solicitudPedidoPyme
	1, // 1: protos.protos.solicitarPedidoRetail:input_type -> protos.solicitudPedidoRetail
	3, // 2: protos.protos.obtenerCodigoSeguimiento:input_type -> protos.solicitudSeguimiento
	5, // 3: protos.protos.solicitarPaquete:input_type -> protos.solicitudPaquete
	7, // 4: protos.protos.obtenerEstado:input_type -> protos.solicitudEstado
	2, // 5: protos.protos.solicitarPedidoPyme:output_type -> protos.respuestaPedido
	2, // 6: protos.protos.solicitarPedidoRetail:output_type -> protos.respuestaPedido
	4, // 7: protos.protos.obtenerCodigoSeguimiento:output_type -> protos.respuestaSeguimiento
	6, // 8: protos.protos.solicitarPaquete:output_type -> protos.respuestaPaquete
	8, // 9: protos.protos.obtenerEstado:output_type -> protos.respuestaEstado
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_proto_init() }
func file_protos_proto_init() {
	if File_protos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudPedidoPyme); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudPedidoRetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaPedido); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudSeguimiento); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaSeguimiento); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudPaquete); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaPaquete); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudEstado); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaEstado); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_proto_goTypes,
		DependencyIndexes: file_protos_proto_depIdxs,
		MessageInfos:      file_protos_proto_msgTypes,
	}.Build()
	File_protos_proto = out.File
	file_protos_proto_rawDesc = nil
	file_protos_proto_goTypes = nil
	file_protos_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtosClient is the client API for Protos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtosClient interface {
	//funciones
	SolicitarPedidoPyme(ctx context.Context, in *SolicitudPedidoPyme, opts ...grpc.CallOption) (*RespuestaPedido, error)
	SolicitarPedidoRetail(ctx context.Context, in *SolicitudPedidoRetail, opts ...grpc.CallOption) (*RespuestaPedido, error)
	ObtenerCodigoSeguimiento(ctx context.Context, in *SolicitudSeguimiento, opts ...grpc.CallOption) (*RespuestaSeguimiento, error)
	SolicitarPaquete(ctx context.Context, in *SolicitudPaquete, opts ...grpc.CallOption) (*RespuestaPaquete, error)
	ObtenerEstado(ctx context.Context, in *SolicitudEstado, opts ...grpc.CallOption) (*RespuestaEstado, error)
}

type protosClient struct {
	cc grpc.ClientConnInterface
}

func NewProtosClient(cc grpc.ClientConnInterface) ProtosClient {
	return &protosClient{cc}
}

func (c *protosClient) SolicitarPedidoPyme(ctx context.Context, in *SolicitudPedidoPyme, opts ...grpc.CallOption) (*RespuestaPedido, error) {
	out := new(RespuestaPedido)
	err := c.cc.Invoke(ctx, "/protos.protos/solicitarPedidoPyme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protosClient) SolicitarPedidoRetail(ctx context.Context, in *SolicitudPedidoRetail, opts ...grpc.CallOption) (*RespuestaPedido, error) {
	out := new(RespuestaPedido)
	err := c.cc.Invoke(ctx, "/protos.protos/solicitarPedidoRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protosClient) ObtenerCodigoSeguimiento(ctx context.Context, in *SolicitudSeguimiento, opts ...grpc.CallOption) (*RespuestaSeguimiento, error) {
	out := new(RespuestaSeguimiento)
	err := c.cc.Invoke(ctx, "/protos.protos/obtenerCodigoSeguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protosClient) SolicitarPaquete(ctx context.Context, in *SolicitudPaquete, opts ...grpc.CallOption) (*RespuestaPaquete, error) {
	out := new(RespuestaPaquete)
	err := c.cc.Invoke(ctx, "/protos.protos/solicitarPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protosClient) ObtenerEstado(ctx context.Context, in *SolicitudEstado, opts ...grpc.CallOption) (*RespuestaEstado, error) {
	out := new(RespuestaEstado)
	err := c.cc.Invoke(ctx, "/protos.protos/obtenerEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtosServer is the server API for Protos service.
type ProtosServer interface {
	//funciones
	SolicitarPedidoPyme(context.Context, *SolicitudPedidoPyme) (*RespuestaPedido, error)
	SolicitarPedidoRetail(context.Context, *SolicitudPedidoRetail) (*RespuestaPedido, error)
	ObtenerCodigoSeguimiento(context.Context, *SolicitudSeguimiento) (*RespuestaSeguimiento, error)
	SolicitarPaquete(context.Context, *SolicitudPaquete) (*RespuestaPaquete, error)
	ObtenerEstado(context.Context, *SolicitudEstado) (*RespuestaEstado, error)
}

// UnimplementedProtosServer can be embedded to have forward compatible implementations.
type UnimplementedProtosServer struct {
}

func (*UnimplementedProtosServer) SolicitarPedidoPyme(context.Context, *SolicitudPedidoPyme) (*RespuestaPedido, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarPedidoPyme not implemented")
}
func (*UnimplementedProtosServer) SolicitarPedidoRetail(context.Context, *SolicitudPedidoRetail) (*RespuestaPedido, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarPedidoRetail not implemented")
}
func (*UnimplementedProtosServer) ObtenerCodigoSeguimiento(context.Context, *SolicitudSeguimiento) (*RespuestaSeguimiento, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerCodigoSeguimiento not implemented")
}
func (*UnimplementedProtosServer) SolicitarPaquete(context.Context, *SolicitudPaquete) (*RespuestaPaquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarPaquete not implemented")
}
func (*UnimplementedProtosServer) ObtenerEstado(context.Context, *SolicitudEstado) (*RespuestaEstado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerEstado not implemented")
}

func RegisterProtosServer(s *grpc.Server, srv ProtosServer) {
	s.RegisterService(&_Protos_serviceDesc, srv)
}

func _Protos_SolicitarPedidoPyme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudPedidoPyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtosServer).SolicitarPedidoPyme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.protos/SolicitarPedidoPyme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtosServer).SolicitarPedidoPyme(ctx, req.(*SolicitudPedidoPyme))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protos_SolicitarPedidoRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudPedidoRetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtosServer).SolicitarPedidoRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.protos/SolicitarPedidoRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtosServer).SolicitarPedidoRetail(ctx, req.(*SolicitudPedidoRetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protos_ObtenerCodigoSeguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudSeguimiento)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtosServer).ObtenerCodigoSeguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.protos/ObtenerCodigoSeguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtosServer).ObtenerCodigoSeguimiento(ctx, req.(*SolicitudSeguimiento))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protos_SolicitarPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudPaquete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtosServer).SolicitarPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.protos/SolicitarPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtosServer).SolicitarPaquete(ctx, req.(*SolicitudPaquete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protos_ObtenerEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudEstado)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtosServer).ObtenerEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.protos/ObtenerEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtosServer).ObtenerEstado(ctx, req.(*SolicitudEstado))
	}
	return interceptor(ctx, in, info, handler)
}

var _Protos_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.protos",
	HandlerType: (*ProtosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "solicitarPedidoPyme",
			Handler:    _Protos_SolicitarPedidoPyme_Handler,
		},
		{
			MethodName: "solicitarPedidoRetail",
			Handler:    _Protos_SolicitarPedidoRetail_Handler,
		},
		{
			MethodName: "obtenerCodigoSeguimiento",
			Handler:    _Protos_ObtenerCodigoSeguimiento_Handler,
		},
		{
			MethodName: "solicitarPaquete",
			Handler:    _Protos_SolicitarPaquete_Handler,
		},
		{
			MethodName: "obtenerEstado",
			Handler:    _Protos_ObtenerEstado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos.proto",
}
