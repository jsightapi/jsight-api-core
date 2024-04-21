package jerr

const (
	RuntimeFailure = "runtime failure" // should not occur
	TestFakeError  = "fake error"

	HTTPResourceNotFound    = "resource not found"
	HTTPMethodNotFound      = "HTTP method not found"
	PathNotFound            = "path not found"
	TagNotFound             = "tag not found"
	UserTypeNotFound        = "user type not found"
	ParentNotFound          = "parent directive not found"
	MacroNotFound           = "macro not found"
	ServerNotFound          = "server not found"
	JsonRpcMethodNotFound   = "JSON-RPC method not found"
	JsonRpcResourceNotFound = "resource not found"

	ProtocolNotFound     = `the directive "Protocol" not found`
	ProtocolParameterErr = `the parameter value have to be "json-rpc-2.0"`

	InfoIsEmpty        = "info is empty"
	ResponsesIsEmpty   = "responses is empty"
	RequestIsEmpty     = "request is empty"
	DescriptionIsEmpty = "empty is description"
	BodyIsEmpty        = "empty is body"
	MacroIsEmpty       = "macro is empty"

	IncorrectPath             = "incorrect path"
	IncorrectRequest          = "incorrect request"
	IncorrectDirectiveContext = "incorrect context for the directive"
	IncorrectParameter        = "incorrect parameter"

	PathOrErr                          = "the root schema object cannot have the `or` rule in the Path directive"
	PathObjectErr                      = "the body of the Path directive must be an object, learn more about the Path directive here: https://jsight.io/docs/jsight-api-0-3#directive-path"
	PathAdditionalPropertiesErr        = `the "additionalProperties" rule is invalid in the Path directive, learn more about the Path directive here: https://jsight.io/docs/jsight-api-0-3#directive-path`
	PathNullableErr                    = `the "nullable" rule is invalid in the Path directive, learn more about the Path directive here: https://jsight.io/docs/jsight-api-0-3#directive-path`
	PathEmptyErr                       = "the object in the Path directive can not be empty, learn more about the Path directive here: https://jsight.io/docs/jsight-api-0-3#directive-path"
	PathMultiLevelPropertyErr          = "the multi-level property is not allowed in the Path directive"
	PathEmptyParameter                 = "empty PATH parameter"
	PathParameterIsDuplicatedInThePath = "the parameter of the path is duplicated"
	PathsAreSimilar                    = "the ambiguous paths are not allowed: \"/%s/{%s}\", \"/%s\", see the details here: https://jsight.io/docs/jsight-api-0-3#parameter-path"

	IncludeRootErr      = "cannot not start with `/`"
	IncludeUpErr        = "cannot contain `..` or `.`"
	IncludeSeparatorErr = "use the symbol `/` as a separator for directories"
	IncludeDirectiveErr = "the directive is not allowed in included files:"

	UnsupportedVersion                = "unsupported version of JSIGHT"
	DirectiveJSIGHTShouldBeTheFirst   = "directive JSIGHT should be the first"
	DirectiveJSIGHTGottaBeOnlyOneTime = "directive JSIGHT gotta be only one time"
	DirectiveINFOGottaBeOnlyOneTime   = "directive INFO gotta be only one time"
	DirectiveBaseURLAlreadyDefined    = "directive BaseURL already defined"

	UnknownDirective = "unknown directive"
	UnknownNotation  = "unknown notation"

	RequiredParameterNotSpecified         = "required parameter(s) not specified"
	ParametersAreForbiddenForTheDirective = "parameters are forbidden for the directive"
	ParametersIsAlreadyDefined            = "the parameter is already defined for the directive"

	AnnotationIsForbiddenForTheDirective                = "annotation is forbidden for the directive"
	NotUniqueDirective                                  = "not a unique directive"
	NotUniquePath                                       = "non-unique path %q in the URL directive"
	BodyMustBeObject                                    = "body must be object"
	CannotUseTheTypeAndSchemaNotationParametersTogether = "cannot use the Type and SchemaNotation parameters together"
	ThereIsNoExplicitContextForClosure                  = "there is no explicit context for closure"
	DirectiveNotAllowed                                 = "directive not allowed"
	ApartFromTheOpeningParenthesis                      = "apart from the opening parenthesis, there should be nothing else on this line" //nolint:lll
	DuplicateNames                                      = "duplicate names are not allowed"
	NotAllowedToOverrideTheProperty                     = "it is not allowed to override the %q property from the user type %q" //nolint:lll
	ContextNotClosed                                    = "not all explicit contexts are closed"
	WrongDescriptionContext                             = "wrong description context"
	MethodIsAlreadyDefinedInResource                    = "method is already defined in resource"
	UndefinedRequestBodyForResource                     = "undefined request body for resource"
	RecursionIsProhibited                               = "recursion is prohibited"
	UserTypeIsNotAnObject                               = "the user type is not an object"
	ProcessTypeErr                                      = "process type"
	FailedToComputeScannersHash                         = "failed to compute scanner's hash"
)
