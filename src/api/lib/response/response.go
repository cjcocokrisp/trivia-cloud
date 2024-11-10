package response

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Definitions for reponses for the websocket api
type Response = events.APIGatewayProxyResponse

func OkReponse() Response {
	return Response{StatusCode: http.StatusOK}
}

func InternalSeverErrorResponse() Response {
	return Response{StatusCode: http.StatusInternalServerError}
}

func BadRequestResponse() Response {
	return Response{StatusCode: http.StatusBadRequest}
}
