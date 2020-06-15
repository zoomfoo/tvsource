// Package runtime provides the Chrome DevTools Protocol
// commands, types, and events for the Runtime domain.
//
// Runtime domain exposes JavaScript runtime by means of remote evaluation
// and mirror objects. Evaluation results are returned as mirror object that
// expose object type, string representation and unique identifier that can be
// used for further object reference. Original objects are maintained in memory
// unless they are either explicitly released or are released along with the
// other objects in their object group.
//
// Generated by the cdproto-gen command.
package runtime

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// AwaitPromiseParams add handler to promise with given promise object id.
type AwaitPromiseParams struct {
	PromiseObjectID RemoteObjectID `json:"promiseObjectId"`           // Identifier of the promise.
	ReturnByValue   bool           `json:"returnByValue,omitempty"`   // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview bool           `json:"generatePreview,omitempty"` // Whether preview should be generated for the result.
}

// AwaitPromise add handler to promise with given promise object id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-awaitPromise
//
// parameters:
//   promiseObjectID - Identifier of the promise.
func AwaitPromise(promiseObjectID RemoteObjectID) *AwaitPromiseParams {
	return &AwaitPromiseParams{
		PromiseObjectID: promiseObjectID,
	}
}

// WithReturnByValue whether the result is expected to be a JSON object that
// should be sent by value.
func (p AwaitPromiseParams) WithReturnByValue(returnByValue bool) *AwaitPromiseParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p AwaitPromiseParams) WithGeneratePreview(generatePreview bool) *AwaitPromiseParams {
	p.GeneratePreview = generatePreview
	return &p
}

// AwaitPromiseReturns return values.
type AwaitPromiseReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Promise result. Will contain rejected value if promise was rejected.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details if stack strace is available.
}

// Do executes Runtime.awaitPromise against the provided context.
//
// returns:
//   result - Promise result. Will contain rejected value if promise was rejected.
//   exceptionDetails - Exception details if stack strace is available.
func (p *AwaitPromiseParams) Do(ctx context.Context) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	// execute
	var res AwaitPromiseReturns
	err = cdp.Execute(ctx, CommandAwaitPromise, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.Result, res.ExceptionDetails, nil
}

// CallFunctionOnParams calls function with given declaration on the given
// object. Object group of the result is inherited from the target object.
type CallFunctionOnParams struct {
	FunctionDeclaration string             `json:"functionDeclaration"`          // Declaration of the function to call.
	ObjectID            RemoteObjectID     `json:"objectId,omitempty"`           // Identifier of the object to call function on. Either objectId or executionContextId should be specified.
	Arguments           []*CallArgument    `json:"arguments,omitempty"`          // Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Silent              bool               `json:"silent,omitempty"`             // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	ReturnByValue       bool               `json:"returnByValue,omitempty"`      // Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview     bool               `json:"generatePreview,omitempty"`    // Whether preview should be generated for the result.
	UserGesture         bool               `json:"userGesture,omitempty"`        // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise        bool               `json:"awaitPromise,omitempty"`       // Whether execution should await for resulting value and return once awaited promise is resolved.
	ExecutionContextID  ExecutionContextID `json:"executionContextId,omitempty"` // Specifies execution context which global object will be used to call function on. Either executionContextId or objectId should be specified.
	ObjectGroup         string             `json:"objectGroup,omitempty"`        // Symbolic group name that can be used to release multiple objects. If objectGroup is not specified and objectId is, objectGroup will be inherited from object.
}

// CallFunctionOn calls function with given declaration on the given object.
// Object group of the result is inherited from the target object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-callFunctionOn
//
// parameters:
//   functionDeclaration - Declaration of the function to call.
func CallFunctionOn(functionDeclaration string) *CallFunctionOnParams {
	return &CallFunctionOnParams{
		FunctionDeclaration: functionDeclaration,
	}
}

// WithObjectID identifier of the object to call function on. Either objectId
// or executionContextId should be specified.
func (p CallFunctionOnParams) WithObjectID(objectID RemoteObjectID) *CallFunctionOnParams {
	p.ObjectID = objectID
	return &p
}

// WithArguments call arguments. All call arguments must belong to the same
// JavaScript world as the target object.
func (p CallFunctionOnParams) WithArguments(arguments []*CallArgument) *CallFunctionOnParams {
	p.Arguments = arguments
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p CallFunctionOnParams) WithSilent(silent bool) *CallFunctionOnParams {
	p.Silent = silent
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object which
// should be sent by value.
func (p CallFunctionOnParams) WithReturnByValue(returnByValue bool) *CallFunctionOnParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p CallFunctionOnParams) WithGeneratePreview(generatePreview bool) *CallFunctionOnParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithUserGesture whether execution should be treated as initiated by user
// in the UI.
func (p CallFunctionOnParams) WithUserGesture(userGesture bool) *CallFunctionOnParams {
	p.UserGesture = userGesture
	return &p
}

// WithAwaitPromise whether execution should await for resulting value and
// return once awaited promise is resolved.
func (p CallFunctionOnParams) WithAwaitPromise(awaitPromise bool) *CallFunctionOnParams {
	p.AwaitPromise = awaitPromise
	return &p
}

// WithExecutionContextID specifies execution context which global object
// will be used to call function on. Either executionContextId or objectId
// should be specified.
func (p CallFunctionOnParams) WithExecutionContextID(executionContextID ExecutionContextID) *CallFunctionOnParams {
	p.ExecutionContextID = executionContextID
	return &p
}

// WithObjectGroup symbolic group name that can be used to release multiple
// objects. If objectGroup is not specified and objectId is, objectGroup will be
// inherited from object.
func (p CallFunctionOnParams) WithObjectGroup(objectGroup string) *CallFunctionOnParams {
	p.ObjectGroup = objectGroup
	return &p
}

// CallFunctionOnReturns return values.
type CallFunctionOnReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Call result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.callFunctionOn against the provided context.
//
// returns:
//   result - Call result.
//   exceptionDetails - Exception details.
func (p *CallFunctionOnParams) Do(ctx context.Context) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	// execute
	var res CallFunctionOnReturns
	err = cdp.Execute(ctx, CommandCallFunctionOn, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.Result, res.ExceptionDetails, nil
}

// CompileScriptParams compiles expression.
type CompileScriptParams struct {
	Expression         string             `json:"expression"`                   // Expression to compile.
	SourceURL          string             `json:"sourceURL"`                    // Source url to be set for the script.
	PersistScript      bool               `json:"persistScript"`                // Specifies whether the compiled script should be persisted.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"` // Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
}

// CompileScript compiles expression.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-compileScript
//
// parameters:
//   expression - Expression to compile.
//   sourceURL - Source url to be set for the script.
//   persistScript - Specifies whether the compiled script should be persisted.
func CompileScript(expression string, sourceURL string, persistScript bool) *CompileScriptParams {
	return &CompileScriptParams{
		Expression:    expression,
		SourceURL:     sourceURL,
		PersistScript: persistScript,
	}
}

// WithExecutionContextID specifies in which execution context to perform
// script run. If the parameter is omitted the evaluation will be performed in
// the context of the inspected page.
func (p CompileScriptParams) WithExecutionContextID(executionContextID ExecutionContextID) *CompileScriptParams {
	p.ExecutionContextID = executionContextID
	return &p
}

// CompileScriptReturns return values.
type CompileScriptReturns struct {
	ScriptID         ScriptID          `json:"scriptId,omitempty"`         // Id of the script.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.compileScript against the provided context.
//
// returns:
//   scriptID - Id of the script.
//   exceptionDetails - Exception details.
func (p *CompileScriptParams) Do(ctx context.Context) (scriptID ScriptID, exceptionDetails *ExceptionDetails, err error) {
	// execute
	var res CompileScriptReturns
	err = cdp.Execute(ctx, CommandCompileScript, p, &res)
	if err != nil {
		return "", nil, err
	}

	return res.ScriptID, res.ExceptionDetails, nil
}

// DisableParams disables reporting of execution contexts creation.
type DisableParams struct{}

// Disable disables reporting of execution contexts creation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Runtime.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// DiscardConsoleEntriesParams discards collected exceptions and console API
// calls.
type DiscardConsoleEntriesParams struct{}

// DiscardConsoleEntries discards collected exceptions and console API calls.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-discardConsoleEntries
func DiscardConsoleEntries() *DiscardConsoleEntriesParams {
	return &DiscardConsoleEntriesParams{}
}

// Do executes Runtime.discardConsoleEntries against the provided context.
func (p *DiscardConsoleEntriesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDiscardConsoleEntries, nil, nil)
}

// EnableParams enables reporting of execution contexts creation by means of
// executionContextCreated event. When the reporting gets enabled the event will
// be sent immediately for each existing execution context.
type EnableParams struct{}

// Enable enables reporting of execution contexts creation by means of
// executionContextCreated event. When the reporting gets enabled the event will
// be sent immediately for each existing execution context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Runtime.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// EvaluateParams evaluates expression on global object.
type EvaluateParams struct {
	Expression            string             `json:"expression"`                      // Expression to evaluate.
	ObjectGroup           string             `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	Silent                bool               `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	ContextID             ExecutionContextID `json:"contextId,omitempty"`             // Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ReturnByValue         bool               `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview       bool               `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	UserGesture           bool               `json:"userGesture,omitempty"`           // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise          bool               `json:"awaitPromise,omitempty"`          // Whether execution should await for resulting value and return once awaited promise is resolved.
	ThrowOnSideEffect     bool               `json:"throwOnSideEffect,omitempty"`     // Whether to throw an exception if side effect cannot be ruled out during evaluation. This implies disableBreaks below.
	Timeout               TimeDelta          `json:"timeout,omitempty"`               // Terminate execution after timing out (number of milliseconds).
	DisableBreaks         bool               `json:"disableBreaks,omitempty"`         // Disable breakpoints during execution.
	ReplMode              bool               `json:"replMode,omitempty"`              // Setting this flag to true enables let re-declaration and top-level await. Note that let variables can only be re-declared if they originate from replMode themselves.
}

// Evaluate evaluates expression on global object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-evaluate
//
// parameters:
//   expression - Expression to evaluate.
func Evaluate(expression string) *EvaluateParams {
	return &EvaluateParams{
		Expression: expression,
	}
}

// WithObjectGroup symbolic group name that can be used to release multiple
// objects.
func (p EvaluateParams) WithObjectGroup(objectGroup string) *EvaluateParams {
	p.ObjectGroup = objectGroup
	return &p
}

// WithIncludeCommandLineAPI determines whether Command Line API should be
// available during the evaluation.
func (p EvaluateParams) WithIncludeCommandLineAPI(includeCommandLineAPI bool) *EvaluateParams {
	p.IncludeCommandLineAPI = includeCommandLineAPI
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p EvaluateParams) WithSilent(silent bool) *EvaluateParams {
	p.Silent = silent
	return &p
}

// WithContextID specifies in which execution context to perform evaluation.
// If the parameter is omitted the evaluation will be performed in the context
// of the inspected page.
func (p EvaluateParams) WithContextID(contextID ExecutionContextID) *EvaluateParams {
	p.ContextID = contextID
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object that
// should be sent by value.
func (p EvaluateParams) WithReturnByValue(returnByValue bool) *EvaluateParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p EvaluateParams) WithGeneratePreview(generatePreview bool) *EvaluateParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithUserGesture whether execution should be treated as initiated by user
// in the UI.
func (p EvaluateParams) WithUserGesture(userGesture bool) *EvaluateParams {
	p.UserGesture = userGesture
	return &p
}

// WithAwaitPromise whether execution should await for resulting value and
// return once awaited promise is resolved.
func (p EvaluateParams) WithAwaitPromise(awaitPromise bool) *EvaluateParams {
	p.AwaitPromise = awaitPromise
	return &p
}

// WithThrowOnSideEffect whether to throw an exception if side effect cannot
// be ruled out during evaluation. This implies disableBreaks below.
func (p EvaluateParams) WithThrowOnSideEffect(throwOnSideEffect bool) *EvaluateParams {
	p.ThrowOnSideEffect = throwOnSideEffect
	return &p
}

// WithTimeout terminate execution after timing out (number of milliseconds).
func (p EvaluateParams) WithTimeout(timeout TimeDelta) *EvaluateParams {
	p.Timeout = timeout
	return &p
}

// WithDisableBreaks disable breakpoints during execution.
func (p EvaluateParams) WithDisableBreaks(disableBreaks bool) *EvaluateParams {
	p.DisableBreaks = disableBreaks
	return &p
}

// WithReplMode setting this flag to true enables let re-declaration and
// top-level await. Note that let variables can only be re-declared if they
// originate from replMode themselves.
func (p EvaluateParams) WithReplMode(replMode bool) *EvaluateParams {
	p.ReplMode = replMode
	return &p
}

// EvaluateReturns return values.
type EvaluateReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Evaluation result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.evaluate against the provided context.
//
// returns:
//   result - Evaluation result.
//   exceptionDetails - Exception details.
func (p *EvaluateParams) Do(ctx context.Context) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	// execute
	var res EvaluateReturns
	err = cdp.Execute(ctx, CommandEvaluate, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.Result, res.ExceptionDetails, nil
}

// GetIsolateIDParams returns the isolate id.
type GetIsolateIDParams struct{}

// GetIsolateID returns the isolate id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-getIsolateId
func GetIsolateID() *GetIsolateIDParams {
	return &GetIsolateIDParams{}
}

// GetIsolateIDReturns return values.
type GetIsolateIDReturns struct {
	ID string `json:"id,omitempty"` // The isolate id.
}

// Do executes Runtime.getIsolateId against the provided context.
//
// returns:
//   id - The isolate id.
func (p *GetIsolateIDParams) Do(ctx context.Context) (id string, err error) {
	// execute
	var res GetIsolateIDReturns
	err = cdp.Execute(ctx, CommandGetIsolateID, nil, &res)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}

// GetHeapUsageParams returns the JavaScript heap usage. It is the total
// usage of the corresponding isolate not scoped to a particular Runtime.
type GetHeapUsageParams struct{}

// GetHeapUsage returns the JavaScript heap usage. It is the total usage of
// the corresponding isolate not scoped to a particular Runtime.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-getHeapUsage
func GetHeapUsage() *GetHeapUsageParams {
	return &GetHeapUsageParams{}
}

// GetHeapUsageReturns return values.
type GetHeapUsageReturns struct {
	UsedSize  float64 `json:"usedSize,omitempty"`  // Used heap size in bytes.
	TotalSize float64 `json:"totalSize,omitempty"` // Allocated heap size in bytes.
}

// Do executes Runtime.getHeapUsage against the provided context.
//
// returns:
//   usedSize - Used heap size in bytes.
//   totalSize - Allocated heap size in bytes.
func (p *GetHeapUsageParams) Do(ctx context.Context) (usedSize float64, totalSize float64, err error) {
	// execute
	var res GetHeapUsageReturns
	err = cdp.Execute(ctx, CommandGetHeapUsage, nil, &res)
	if err != nil {
		return 0, 0, err
	}

	return res.UsedSize, res.TotalSize, nil
}

// GetPropertiesParams returns properties of a given object. Object group of
// the result is inherited from the target object.
type GetPropertiesParams struct {
	ObjectID               RemoteObjectID `json:"objectId"`                         // Identifier of the object to return properties for.
	OwnProperties          bool           `json:"ownProperties,omitempty"`          // If true, returns properties belonging only to the element itself, not to its prototype chain.
	AccessorPropertiesOnly bool           `json:"accessorPropertiesOnly,omitempty"` // If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
	GeneratePreview        bool           `json:"generatePreview,omitempty"`        // Whether preview should be generated for the results.
}

// GetProperties returns properties of a given object. Object group of the
// result is inherited from the target object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-getProperties
//
// parameters:
//   objectID - Identifier of the object to return properties for.
func GetProperties(objectID RemoteObjectID) *GetPropertiesParams {
	return &GetPropertiesParams{
		ObjectID: objectID,
	}
}

// WithOwnProperties if true, returns properties belonging only to the
// element itself, not to its prototype chain.
func (p GetPropertiesParams) WithOwnProperties(ownProperties bool) *GetPropertiesParams {
	p.OwnProperties = ownProperties
	return &p
}

// WithAccessorPropertiesOnly if true, returns accessor properties (with
// getter/setter) only; internal properties are not returned either.
func (p GetPropertiesParams) WithAccessorPropertiesOnly(accessorPropertiesOnly bool) *GetPropertiesParams {
	p.AccessorPropertiesOnly = accessorPropertiesOnly
	return &p
}

// WithGeneratePreview whether preview should be generated for the results.
func (p GetPropertiesParams) WithGeneratePreview(generatePreview bool) *GetPropertiesParams {
	p.GeneratePreview = generatePreview
	return &p
}

// GetPropertiesReturns return values.
type GetPropertiesReturns struct {
	Result             []*PropertyDescriptor         `json:"result,omitempty"`             // Object properties.
	InternalProperties []*InternalPropertyDescriptor `json:"internalProperties,omitempty"` // Internal object properties (only of the element itself).
	PrivateProperties  []*PrivatePropertyDescriptor  `json:"privateProperties,omitempty"`  // Object private properties.
	ExceptionDetails   *ExceptionDetails             `json:"exceptionDetails,omitempty"`   // Exception details.
}

// Do executes Runtime.getProperties against the provided context.
//
// returns:
//   result - Object properties.
//   internalProperties - Internal object properties (only of the element itself).
//   privateProperties - Object private properties.
//   exceptionDetails - Exception details.
func (p *GetPropertiesParams) Do(ctx context.Context) (result []*PropertyDescriptor, internalProperties []*InternalPropertyDescriptor, privateProperties []*PrivatePropertyDescriptor, exceptionDetails *ExceptionDetails, err error) {
	// execute
	var res GetPropertiesReturns
	err = cdp.Execute(ctx, CommandGetProperties, p, &res)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return res.Result, res.InternalProperties, res.PrivateProperties, res.ExceptionDetails, nil
}

// GlobalLexicalScopeNamesParams returns all let, const and class variables
// from global scope.
type GlobalLexicalScopeNamesParams struct {
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"` // Specifies in which execution context to lookup global scope variables.
}

// GlobalLexicalScopeNames returns all let, const and class variables from
// global scope.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-globalLexicalScopeNames
//
// parameters:
func GlobalLexicalScopeNames() *GlobalLexicalScopeNamesParams {
	return &GlobalLexicalScopeNamesParams{}
}

// WithExecutionContextID specifies in which execution context to lookup
// global scope variables.
func (p GlobalLexicalScopeNamesParams) WithExecutionContextID(executionContextID ExecutionContextID) *GlobalLexicalScopeNamesParams {
	p.ExecutionContextID = executionContextID
	return &p
}

// GlobalLexicalScopeNamesReturns return values.
type GlobalLexicalScopeNamesReturns struct {
	Names []string `json:"names,omitempty"`
}

// Do executes Runtime.globalLexicalScopeNames against the provided context.
//
// returns:
//   names
func (p *GlobalLexicalScopeNamesParams) Do(ctx context.Context) (names []string, err error) {
	// execute
	var res GlobalLexicalScopeNamesReturns
	err = cdp.Execute(ctx, CommandGlobalLexicalScopeNames, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Names, nil
}

// QueryObjectsParams [no description].
type QueryObjectsParams struct {
	PrototypeObjectID RemoteObjectID `json:"prototypeObjectId"`     // Identifier of the prototype to return objects for.
	ObjectGroup       string         `json:"objectGroup,omitempty"` // Symbolic group name that can be used to release the results.
}

// QueryObjects [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-queryObjects
//
// parameters:
//   prototypeObjectID - Identifier of the prototype to return objects for.
func QueryObjects(prototypeObjectID RemoteObjectID) *QueryObjectsParams {
	return &QueryObjectsParams{
		PrototypeObjectID: prototypeObjectID,
	}
}

// WithObjectGroup symbolic group name that can be used to release the
// results.
func (p QueryObjectsParams) WithObjectGroup(objectGroup string) *QueryObjectsParams {
	p.ObjectGroup = objectGroup
	return &p
}

// QueryObjectsReturns return values.
type QueryObjectsReturns struct {
	Objects *RemoteObject `json:"objects,omitempty"` // Array with objects.
}

// Do executes Runtime.queryObjects against the provided context.
//
// returns:
//   objects - Array with objects.
func (p *QueryObjectsParams) Do(ctx context.Context) (objects *RemoteObject, err error) {
	// execute
	var res QueryObjectsReturns
	err = cdp.Execute(ctx, CommandQueryObjects, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Objects, nil
}

// ReleaseObjectParams releases remote object with given id.
type ReleaseObjectParams struct {
	ObjectID RemoteObjectID `json:"objectId"` // Identifier of the object to release.
}

// ReleaseObject releases remote object with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-releaseObject
//
// parameters:
//   objectID - Identifier of the object to release.
func ReleaseObject(objectID RemoteObjectID) *ReleaseObjectParams {
	return &ReleaseObjectParams{
		ObjectID: objectID,
	}
}

// Do executes Runtime.releaseObject against the provided context.
func (p *ReleaseObjectParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandReleaseObject, p, nil)
}

// ReleaseObjectGroupParams releases all remote objects that belong to a
// given group.
type ReleaseObjectGroupParams struct {
	ObjectGroup string `json:"objectGroup"` // Symbolic object group name.
}

// ReleaseObjectGroup releases all remote objects that belong to a given
// group.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-releaseObjectGroup
//
// parameters:
//   objectGroup - Symbolic object group name.
func ReleaseObjectGroup(objectGroup string) *ReleaseObjectGroupParams {
	return &ReleaseObjectGroupParams{
		ObjectGroup: objectGroup,
	}
}

// Do executes Runtime.releaseObjectGroup against the provided context.
func (p *ReleaseObjectGroupParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandReleaseObjectGroup, p, nil)
}

// RunIfWaitingForDebuggerParams tells inspected instance to run if it was
// waiting for debugger to attach.
type RunIfWaitingForDebuggerParams struct{}

// RunIfWaitingForDebugger tells inspected instance to run if it was waiting
// for debugger to attach.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-runIfWaitingForDebugger
func RunIfWaitingForDebugger() *RunIfWaitingForDebuggerParams {
	return &RunIfWaitingForDebuggerParams{}
}

// Do executes Runtime.runIfWaitingForDebugger against the provided context.
func (p *RunIfWaitingForDebuggerParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandRunIfWaitingForDebugger, nil, nil)
}

// RunScriptParams runs script with given id in a given context.
type RunScriptParams struct {
	ScriptID              ScriptID           `json:"scriptId"`                        // Id of the script to run.
	ExecutionContextID    ExecutionContextID `json:"executionContextId,omitempty"`    // Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ObjectGroup           string             `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	Silent                bool               `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	ReturnByValue         bool               `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview       bool               `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	AwaitPromise          bool               `json:"awaitPromise,omitempty"`          // Whether execution should await for resulting value and return once awaited promise is resolved.
}

// RunScript runs script with given id in a given context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-runScript
//
// parameters:
//   scriptID - Id of the script to run.
func RunScript(scriptID ScriptID) *RunScriptParams {
	return &RunScriptParams{
		ScriptID: scriptID,
	}
}

// WithExecutionContextID specifies in which execution context to perform
// script run. If the parameter is omitted the evaluation will be performed in
// the context of the inspected page.
func (p RunScriptParams) WithExecutionContextID(executionContextID ExecutionContextID) *RunScriptParams {
	p.ExecutionContextID = executionContextID
	return &p
}

// WithObjectGroup symbolic group name that can be used to release multiple
// objects.
func (p RunScriptParams) WithObjectGroup(objectGroup string) *RunScriptParams {
	p.ObjectGroup = objectGroup
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p RunScriptParams) WithSilent(silent bool) *RunScriptParams {
	p.Silent = silent
	return &p
}

// WithIncludeCommandLineAPI determines whether Command Line API should be
// available during the evaluation.
func (p RunScriptParams) WithIncludeCommandLineAPI(includeCommandLineAPI bool) *RunScriptParams {
	p.IncludeCommandLineAPI = includeCommandLineAPI
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object which
// should be sent by value.
func (p RunScriptParams) WithReturnByValue(returnByValue bool) *RunScriptParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p RunScriptParams) WithGeneratePreview(generatePreview bool) *RunScriptParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithAwaitPromise whether execution should await for resulting value and
// return once awaited promise is resolved.
func (p RunScriptParams) WithAwaitPromise(awaitPromise bool) *RunScriptParams {
	p.AwaitPromise = awaitPromise
	return &p
}

// RunScriptReturns return values.
type RunScriptReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Run result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.runScript against the provided context.
//
// returns:
//   result - Run result.
//   exceptionDetails - Exception details.
func (p *RunScriptParams) Do(ctx context.Context) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	// execute
	var res RunScriptReturns
	err = cdp.Execute(ctx, CommandRunScript, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.Result, res.ExceptionDetails, nil
}

// SetCustomObjectFormatterEnabledParams [no description].
type SetCustomObjectFormatterEnabledParams struct {
	Enabled bool `json:"enabled"`
}

// SetCustomObjectFormatterEnabled [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-setCustomObjectFormatterEnabled
//
// parameters:
//   enabled
func SetCustomObjectFormatterEnabled(enabled bool) *SetCustomObjectFormatterEnabledParams {
	return &SetCustomObjectFormatterEnabledParams{
		Enabled: enabled,
	}
}

// Do executes Runtime.setCustomObjectFormatterEnabled against the provided context.
func (p *SetCustomObjectFormatterEnabledParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetCustomObjectFormatterEnabled, p, nil)
}

// SetMaxCallStackSizeToCaptureParams [no description].
type SetMaxCallStackSizeToCaptureParams struct {
	Size int64 `json:"size"`
}

// SetMaxCallStackSizeToCapture [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-setMaxCallStackSizeToCapture
//
// parameters:
//   size
func SetMaxCallStackSizeToCapture(size int64) *SetMaxCallStackSizeToCaptureParams {
	return &SetMaxCallStackSizeToCaptureParams{
		Size: size,
	}
}

// Do executes Runtime.setMaxCallStackSizeToCapture against the provided context.
func (p *SetMaxCallStackSizeToCaptureParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetMaxCallStackSizeToCapture, p, nil)
}

// TerminateExecutionParams terminate current or next JavaScript execution.
// Will cancel the termination when the outer-most script execution ends.
type TerminateExecutionParams struct{}

// TerminateExecution terminate current or next JavaScript execution. Will
// cancel the termination when the outer-most script execution ends.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-terminateExecution
func TerminateExecution() *TerminateExecutionParams {
	return &TerminateExecutionParams{}
}

// Do executes Runtime.terminateExecution against the provided context.
func (p *TerminateExecutionParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTerminateExecution, nil, nil)
}

// AddBindingParams if executionContextId is empty, adds binding with the
// given name on the global objects of all inspected contexts, including those
// created later, bindings survive reloads. If executionContextId is specified,
// adds binding only on global object of given execution context. Binding
// function takes exactly one argument, this argument should be string, in case
// of any other input, function throws an exception. Each binding function call
// produces Runtime.bindingCalled notification.
type AddBindingParams struct {
	Name               string             `json:"name"`
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

// AddBinding if executionContextId is empty, adds binding with the given
// name on the global objects of all inspected contexts, including those created
// later, bindings survive reloads. If executionContextId is specified, adds
// binding only on global object of given execution context. Binding function
// takes exactly one argument, this argument should be string, in case of any
// other input, function throws an exception. Each binding function call
// produces Runtime.bindingCalled notification.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-addBinding
//
// parameters:
//   name
func AddBinding(name string) *AddBindingParams {
	return &AddBindingParams{
		Name: name,
	}
}

// WithExecutionContextID [no description].
func (p AddBindingParams) WithExecutionContextID(executionContextID ExecutionContextID) *AddBindingParams {
	p.ExecutionContextID = executionContextID
	return &p
}

// Do executes Runtime.addBinding against the provided context.
func (p *AddBindingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandAddBinding, p, nil)
}

// RemoveBindingParams this method does not remove binding function from
// global object but unsubscribes current runtime agent from
// Runtime.bindingCalled notifications.
type RemoveBindingParams struct {
	Name string `json:"name"`
}

// RemoveBinding this method does not remove binding function from global
// object but unsubscribes current runtime agent from Runtime.bindingCalled
// notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#method-removeBinding
//
// parameters:
//   name
func RemoveBinding(name string) *RemoveBindingParams {
	return &RemoveBindingParams{
		Name: name,
	}
}

// Do executes Runtime.removeBinding against the provided context.
func (p *RemoveBindingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandRemoveBinding, p, nil)
}

// Command names.
const (
	CommandAwaitPromise                    = "Runtime.awaitPromise"
	CommandCallFunctionOn                  = "Runtime.callFunctionOn"
	CommandCompileScript                   = "Runtime.compileScript"
	CommandDisable                         = "Runtime.disable"
	CommandDiscardConsoleEntries           = "Runtime.discardConsoleEntries"
	CommandEnable                          = "Runtime.enable"
	CommandEvaluate                        = "Runtime.evaluate"
	CommandGetIsolateID                    = "Runtime.getIsolateId"
	CommandGetHeapUsage                    = "Runtime.getHeapUsage"
	CommandGetProperties                   = "Runtime.getProperties"
	CommandGlobalLexicalScopeNames         = "Runtime.globalLexicalScopeNames"
	CommandQueryObjects                    = "Runtime.queryObjects"
	CommandReleaseObject                   = "Runtime.releaseObject"
	CommandReleaseObjectGroup              = "Runtime.releaseObjectGroup"
	CommandRunIfWaitingForDebugger         = "Runtime.runIfWaitingForDebugger"
	CommandRunScript                       = "Runtime.runScript"
	CommandSetCustomObjectFormatterEnabled = "Runtime.setCustomObjectFormatterEnabled"
	CommandSetMaxCallStackSizeToCapture    = "Runtime.setMaxCallStackSizeToCapture"
	CommandTerminateExecution              = "Runtime.terminateExecution"
	CommandAddBinding                      = "Runtime.addBinding"
	CommandRemoveBinding                   = "Runtime.removeBinding"
)
