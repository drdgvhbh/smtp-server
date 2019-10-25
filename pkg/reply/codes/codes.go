package codes

type ResponseCode uint

const (
	UnableToConnect                  ResponseCode = 101
	ConnectionRefused                ResponseCode = 111
	SystemStatusOrHelpReply          ResponseCode = 200
	HelpResponse                     ResponseCode = 214
	Ready                            ResponseCode = 220
	ClosingTransmissionChannel       ResponseCode = 221
	Ok                               ResponseCode = 250
	UserNotLocalWilLForward          ResponseCode = 251
	CannotVerifyUser                 ResponseCode = 252
	StartMailInput                   ResponseCode = 354
	Timeout                          ResponseCode = 420
	ServiceUnavailable               ResponseCode = 421
	MailboxLimitExceeded             ResponseCode = 422
	NotEnoughSpaceOnDisk             ResponseCode = 431
	IncomingMailQueueStopped         ResponseCode = 432
	ServerNotResponding              ResponseCode = 441
	ConnectionDropped                ResponseCode = 442
	MaximumHopCountExceeded          ResponseCode = 446
	MessageTimedOut                  ResponseCode = 447
	RoutingError                     ResponseCode = 449
	MailboxUnavailable               ResponseCode = 450
	Aborted                          ResponseCode = 451
	TooManyEmailsOrRecipients        ResponseCode = 452
	MailServerError                  ResponseCode = 471
	SyntaxError                      ResponseCode = 500
	SyntaxErrorInParamsOrArgs        ResponseCode = 501
	BadSequenceOfCommands            ResponseCode = 503
	CommandParameterIsNotImplemented ResponseCode = 504
	BadEmailAddress                  ResponseCode = 510
	BadEmailAddress2                 ResponseCode = 511
	DNSNotFound                      ResponseCode = 512
	AddressTypeIncorrect             ResponseCode = 513
	SizeOfMailExceedsServerLimits    ResponseCode = 523
	AuthenticationProblem            ResponseCode = 530
	RecipientAddressRejectedMsg      ResponseCode = 541
	NonExistentEmailAddress          ResponseCode = 550
	RelayDenied                      ResponseCode = 551
	ExceededStorageAllocation        ResponseCode = 552
	MailboxNameInvalid               ResponseCode = 553
	TransactionHasFailed             ResponseCode = 554
)
