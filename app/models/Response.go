package models


type Response struct  {
	Ok bool
	Msg string

}

func NewResponse() Response  {
	return Response{Ok: false}
}
