package terrors

const defaultStatusCode int = 500

var (
	defaultErrorMessage = externalMessage{
		Description:  "can't get error for your case",
		InternalCode: 399999,
	}

	externalMessagesMap = map[uint32]*serializedExternalMessage{}
)
