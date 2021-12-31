// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BindESUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BindESUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindESUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerResponseBody) SetRequestId(v string) *BindESUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindESUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *BindESUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type BindESUserAnalyzerResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindESUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s BindESUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerResponse) SetHeaders(v map[string]*string) *BindESUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *BindESUserAnalyzerResponse) SetBody(v *BindESUserAnalyzerResponseBody) *BindESUserAnalyzerResponse {
	s.Body = v
	return s
}

type BindEsInstanceResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BindEsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BindEsInstanceResponseBody) SetRequestId(v string) *BindEsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEsInstanceResponseBody) SetResult(v map[string]interface{}) *BindEsInstanceResponseBody {
	s.Result = v
	return s
}

type BindEsInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindEsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEsInstanceResponse) GoString() string {
	return s.String()
}

func (s *BindEsInstanceResponse) SetHeaders(v map[string]*string) *BindEsInstanceResponse {
	s.Headers = v
	return s
}

func (s *BindEsInstanceResponse) SetBody(v *BindEsInstanceResponseBody) *BindEsInstanceResponse {
	s.Body = v
	return s
}

type CompileSortScriptResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CompileSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompileSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CompileSortScriptResponseBody) SetRequestId(v string) *CompileSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type CompileSortScriptResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompileSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompileSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s CompileSortScriptResponse) GoString() string {
	return s.String()
}

func (s *CompileSortScriptResponse) SetHeaders(v map[string]*string) *CompileSortScriptResponse {
	s.Headers = v
	return s
}

func (s *CompileSortScriptResponse) SetBody(v *CompileSortScriptResponseBody) *CompileSortScriptResponse {
	s.Body = v
	return s
}

type CreateABTestExperimentResponseBody struct {
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponseBody) SetRequestId(v string) *CreateABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestExperimentResponseBody) SetResult(v *CreateABTestExperimentResponseBodyResult) *CreateABTestExperimentResponseBody {
	s.Result = v
	return s
}

type CreateABTestExperimentResponseBodyResult struct {
	Created *int32                 `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string                `json:"name,omitempty" xml:"name,omitempty"`
	Online  *bool                  `json:"online,omitempty" xml:"online,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Traffic *int32                 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	Updated *int32                 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateABTestExperimentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponseBodyResult) SetCreated(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetId(v string) *CreateABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetName(v string) *CreateABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetOnline(v bool) *CreateABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetParams(v map[string]interface{}) *CreateABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetTraffic(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetUpdated(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateABTestExperimentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponse) SetHeaders(v map[string]*string) *CreateABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestExperimentResponse) SetBody(v *CreateABTestExperimentResponseBody) *CreateABTestExperimentResponse {
	s.Body = v
	return s
}

type CreateABTestGroupResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponseBody) SetRequestId(v string) *CreateABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestGroupResponseBody) SetResult(v *CreateABTestGroupResponseBodyResult) *CreateABTestGroupResponseBody {
	s.Result = v
	return s
}

type CreateABTestGroupResponseBodyResult struct {
	Created *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateABTestGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponseBodyResult) SetCreated(v int32) *CreateABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetId(v string) *CreateABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetName(v string) *CreateABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetStatus(v int32) *CreateABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetUpdated(v int32) *CreateABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateABTestGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponse) SetHeaders(v map[string]*string) *CreateABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestGroupResponse) SetBody(v *CreateABTestGroupResponseBody) *CreateABTestGroupResponse {
	s.Body = v
	return s
}

type CreateABTestSceneResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponseBody) SetRequestId(v string) *CreateABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestSceneResponseBody) SetResult(v *CreateABTestSceneResponseBodyResult) *CreateABTestSceneResponseBody {
	s.Result = v
	return s
}

type CreateABTestSceneResponseBodyResult struct {
	Created *int32    `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string   `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string   `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32    `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32    `json:"updated,omitempty" xml:"updated,omitempty"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateABTestSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponseBodyResult) SetCreated(v int32) *CreateABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetId(v string) *CreateABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetName(v string) *CreateABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetStatus(v int32) *CreateABTestSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetUpdated(v int32) *CreateABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetValues(v []*string) *CreateABTestSceneResponseBodyResult {
	s.Values = v
	return s
}

type CreateABTestSceneResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponse) SetHeaders(v map[string]*string) *CreateABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestSceneResponse) SetBody(v *CreateABTestSceneResponseBody) *CreateABTestSceneResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetDryRun(v bool) *CreateAppRequest {
	s.DryRun = &v
	return s
}

type CreateAppResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v map[string]interface{}) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateAppGroupResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBody) SetRequestId(v string) *CreateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetResult(v *CreateAppGroupResponseBodyResult) *CreateAppGroupResponseBody {
	s.Result = v
	return s
}

type CreateAppGroupResponseBodyResult struct {
	ChargeType                        *string                                `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay                       *int32                                 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	CommodityCode                     *string                                `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Created                           *int32                                 `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion                    *string                                `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Description                       *string                                `json:"description,omitempty" xml:"description,omitempty"`
	Domain                            *string                                `json:"domain,omitempty" xml:"domain,omitempty"`
	ExpireOn                          *string                                `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	FirstRankAlgoDeploymentId         *int32                                 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	HasPendingQuotaReviewTask         *int32                                 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	Id                                *string                                `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId                        *string                                `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockMode                          *string                                `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration                *int32                                 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	Name                              *string                                `json:"name,omitempty" xml:"name,omitempty"`
	PendingSecondRankAlgoDeploymentId *int32                                 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	ProcessingOrderId                 *string                                `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	Produced                          *int32                                 `json:"produced,omitempty" xml:"produced,omitempty"`
	ProjectId                         *string                                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Quota                             *CreateAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	SecondRankAlgoDeploymentId        *int32                                 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	Status                            *string                                `json:"status,omitempty" xml:"status,omitempty"`
	SwitchedTime                      *int32                                 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Type                              *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Updated                           *int32                                 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResult) SetChargeType(v string) *CreateAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetChargingWay(v int32) *CreateAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCommodityCode(v string) *CreateAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCreated(v int32) *CreateAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCurrentVersion(v string) *CreateAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetDescription(v string) *CreateAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetDomain(v string) *CreateAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetExpireOn(v string) *CreateAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *CreateAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *CreateAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetId(v string) *CreateAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetInstanceId(v string) *CreateAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetLockMode(v string) *CreateAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *CreateAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetName(v string) *CreateAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *CreateAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProcessingOrderId(v string) *CreateAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProduced(v int32) *CreateAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProjectId(v string) *CreateAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetQuota(v *CreateAppGroupResponseBodyResultQuota) *CreateAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *CreateAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetStatus(v string) *CreateAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetSwitchedTime(v int32) *CreateAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetType(v string) *CreateAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetUpdated(v int32) *CreateAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateAppGroupResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateAppGroupResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *CreateAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) SetDocSize(v int32) *CreateAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) SetSpec(v string) *CreateAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type CreateAppGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponse) SetHeaders(v map[string]*string) *CreateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupResponse) SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse {
	s.Body = v
	return s
}

type CreateDataCollectionResponseBody struct {
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateDataCollectionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateDataCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCollectionResponseBody) SetRequestId(v string) *CreateDataCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataCollectionResponseBody) SetResult(v *CreateDataCollectionResponseBodyResult) *CreateDataCollectionResponseBody {
	s.Result = v
	return s
}

type CreateDataCollectionResponseBodyResult struct {
	Created            *int32  `json:"created,omitempty" xml:"created,omitempty"`
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	Id                 *string `json:"id,omitempty" xml:"id,omitempty"`
	IndustryName       *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Status             *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SundialId          *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated            *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateDataCollectionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCollectionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDataCollectionResponseBodyResult) SetCreated(v int32) *CreateDataCollectionResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetDataCollectionType(v string) *CreateDataCollectionResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetId(v string) *CreateDataCollectionResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetIndustryName(v string) *CreateDataCollectionResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetName(v string) *CreateDataCollectionResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetStatus(v int32) *CreateDataCollectionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetSundialId(v string) *CreateDataCollectionResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetType(v string) *CreateDataCollectionResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateDataCollectionResponseBodyResult) SetUpdated(v int32) *CreateDataCollectionResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateDataCollectionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCollectionResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCollectionResponse) SetHeaders(v map[string]*string) *CreateDataCollectionResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCollectionResponse) SetBody(v *CreateDataCollectionResponseBody) *CreateDataCollectionResponse {
	s.Body = v
	return s
}

type CreateFirstRankRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateFirstRankRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankRequest) GoString() string {
	return s.String()
}

func (s *CreateFirstRankRequest) SetDryRun(v bool) *CreateFirstRankRequest {
	s.DryRun = &v
	return s
}

type CreateFirstRankResponseBody struct {
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBody) SetRequestId(v string) *CreateFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirstRankResponseBody) SetResult(v *CreateFirstRankResponseBodyResult) *CreateFirstRankResponseBody {
	s.Result = v
	return s
}

type CreateFirstRankResponseBodyResult struct {
	Active *bool                                    `json:"active,omitempty" xml:"active,omitempty"`
	Meta   []*CreateFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name   *string                                  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBodyResult) SetActive(v bool) *CreateFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *CreateFirstRankResponseBodyResult) SetMeta(v []*CreateFirstRankResponseBodyResultMeta) *CreateFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *CreateFirstRankResponseBodyResult) SetName(v string) *CreateFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type CreateFirstRankResponseBodyResultMeta struct {
	Arg       *string  `json:"arg,omitempty" xml:"arg,omitempty"`
	Attribute *string  `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Weight    *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBodyResultMeta) SetArg(v string) *CreateFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) SetAttribute(v string) *CreateFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) SetWeight(v float32) *CreateFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type CreateFirstRankResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponse) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponse) SetHeaders(v map[string]*string) *CreateFirstRankResponse {
	s.Headers = v
	return s
}

func (s *CreateFirstRankResponse) SetBody(v *CreateFirstRankResponseBody) *CreateFirstRankResponse {
	s.Body = v
	return s
}

type CreateFunctionInstanceRequest struct {
	// 创建参数
	CreateParameters []*CreateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// 周期训练
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// 实例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 功能类型
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// 实例名称
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// 模型类型
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
}

func (s CreateFunctionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequest) SetCreateParameters(v []*CreateFunctionInstanceRequestCreateParameters) *CreateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *CreateFunctionInstanceRequest) SetCron(v string) *CreateFunctionInstanceRequest {
	s.Cron = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetDescription(v string) *CreateFunctionInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetFunctionType(v string) *CreateFunctionInstanceRequest {
	s.FunctionType = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetInstanceName(v string) *CreateFunctionInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetModelType(v string) *CreateFunctionInstanceRequest {
	s.ModelType = &v
	return s
}

type CreateFunctionInstanceRequestCreateParameters struct {
	// 参数名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 参数值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateFunctionInstanceRequestCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequestCreateParameters) SetName(v string) *CreateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *CreateFunctionInstanceRequestCreateParameters) SetValue(v string) *CreateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

type CreateFunctionInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64  `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceResponseBody) SetCode(v string) *CreateFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetHttpCode(v int64) *CreateFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetLatency(v int64) *CreateFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetMessage(v string) *CreateFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetRequestId(v string) *CreateFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetStatus(v string) *CreateFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateFunctionInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceResponse) SetHeaders(v map[string]*string) *CreateFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionInstanceResponse) SetBody(v *CreateFunctionInstanceResponseBody) *CreateFunctionInstanceResponse {
	s.Body = v
	return s
}

type CreateFunctionTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64  `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskResponseBody) SetCode(v string) *CreateFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetHttpCode(v int64) *CreateFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetLatency(v int64) *CreateFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetMessage(v string) *CreateFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetRequestId(v string) *CreateFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetStatus(v string) *CreateFunctionTaskResponseBody {
	s.Status = &v
	return s
}

type CreateFunctionTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskResponse) SetHeaders(v map[string]*string) *CreateFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionTaskResponse) SetBody(v *CreateFunctionTaskResponseBody) *CreateFunctionTaskResponse {
	s.Body = v
	return s
}

type CreateInterventionDictionaryResponseBody struct {
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInterventionDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponseBody) SetRequestId(v string) *CreateInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBody) SetResult(v *CreateInterventionDictionaryResponseBodyResult) *CreateInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

type CreateInterventionDictionaryResponseBodyResult struct {
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	Created  *string `json:"created,omitempty" xml:"created,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated  *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateInterventionDictionaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetCreated(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetName(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetType(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetUpdated(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateInterventionDictionaryResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInterventionDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponse) SetHeaders(v map[string]*string) *CreateInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *CreateInterventionDictionaryResponse) SetBody(v *CreateInterventionDictionaryResponseBody) *CreateInterventionDictionaryResponse {
	s.Body = v
	return s
}

type CreateModelResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetResult(v map[string]interface{}) *CreateModelResponseBody {
	s.Result = v
	return s
}

type CreateModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

type CreateQueryProcessorRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateQueryProcessorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorRequest) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorRequest) SetDryRun(v bool) *CreateQueryProcessorRequest {
	s.DryRun = &v
	return s
}

type CreateQueryProcessorResponseBody struct {
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponseBody) SetRequestId(v string) *CreateQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueryProcessorResponseBody) SetResult(v *CreateQueryProcessorResponseBodyResult) *CreateQueryProcessorResponseBody {
	s.Result = v
	return s
}

type CreateQueryProcessorResponseBodyResult struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Created    *int32                   `json:"created,omitempty" xml:"created,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	Updated    *int32                   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateQueryProcessorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponseBodyResult) SetActive(v bool) *CreateQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetCreated(v int32) *CreateQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetDomain(v string) *CreateQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetIndexes(v []*string) *CreateQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetName(v string) *CreateQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *CreateQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetUpdated(v int32) *CreateQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateQueryProcessorResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponse) SetHeaders(v map[string]*string) *CreateQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *CreateQueryProcessorResponse) SetBody(v *CreateQueryProcessorResponseBody) *CreateQueryProcessorResponse {
	s.Body = v
	return s
}

type CreateScheduledTaskResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetResult(v map[string]interface{}) *CreateScheduledTaskResponseBody {
	s.Result = v
	return s
}

type CreateScheduledTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponse) SetHeaders(v map[string]*string) *CreateScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledTaskResponse) SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse {
	s.Body = v
	return s
}

type CreateSearchStrategyResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyResponseBody) SetRequestId(v string) *CreateSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type CreateSearchStrategyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyResponse) SetHeaders(v map[string]*string) *CreateSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchStrategyResponse) SetBody(v *CreateSearchStrategyResponseBody) *CreateSearchStrategyResponse {
	s.Body = v
	return s
}

type CreateSecondRankRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateSecondRankRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondRankRequest) GoString() string {
	return s.String()
}

func (s *CreateSecondRankRequest) SetDryRun(v bool) *CreateSecondRankRequest {
	s.DryRun = &v
	return s
}

type CreateSecondRankResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateSecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecondRankResponseBody) SetRequestId(v string) *CreateSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecondRankResponseBody) SetResult(v map[string]interface{}) *CreateSecondRankResponseBody {
	s.Result = v
	return s
}

type CreateSecondRankResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondRankResponse) GoString() string {
	return s.String()
}

func (s *CreateSecondRankResponse) SetHeaders(v map[string]*string) *CreateSecondRankResponse {
	s.Headers = v
	return s
}

func (s *CreateSecondRankResponse) SetBody(v *CreateSecondRankResponseBody) *CreateSecondRankResponse {
	s.Body = v
	return s
}

type CreateSortScriptResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponseBody) SetRequestId(v string) *CreateSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type CreateSortScriptResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSortScriptResponse) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponse) SetHeaders(v map[string]*string) *CreateSortScriptResponse {
	s.Headers = v
	return s
}

func (s *CreateSortScriptResponse) SetBody(v *CreateSortScriptResponseBody) *CreateSortScriptResponse {
	s.Body = v
	return s
}

type CreateUserAnalyzerResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerResponseBody) SetRequestId(v string) *CreateUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *CreateUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type CreateUserAnalyzerResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerResponse) SetHeaders(v map[string]*string) *CreateUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *CreateUserAnalyzerResponse) SetBody(v *CreateUserAnalyzerResponseBody) *CreateUserAnalyzerResponse {
	s.Body = v
	return s
}

type DeleteABTestExperimentResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestExperimentResponseBody) SetRequestId(v string) *DeleteABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestExperimentResponseBody) SetResult(v map[string]interface{}) *DeleteABTestExperimentResponseBody {
	s.Result = v
	return s
}

type DeleteABTestExperimentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestExperimentResponse) SetHeaders(v map[string]*string) *DeleteABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestExperimentResponse) SetBody(v *DeleteABTestExperimentResponseBody) *DeleteABTestExperimentResponse {
	s.Body = v
	return s
}

type DeleteABTestGroupResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestGroupResponseBody) SetRequestId(v string) *DeleteABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestGroupResponseBody) SetResult(v map[string]interface{}) *DeleteABTestGroupResponseBody {
	s.Result = v
	return s
}

type DeleteABTestGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestGroupResponse) SetHeaders(v map[string]*string) *DeleteABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestGroupResponse) SetBody(v *DeleteABTestGroupResponseBody) *DeleteABTestGroupResponse {
	s.Body = v
	return s
}

type DeleteABTestSceneResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestSceneResponseBody) SetRequestId(v string) *DeleteABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestSceneResponseBody) SetResult(v map[string]interface{}) *DeleteABTestSceneResponseBody {
	s.Result = v
	return s
}

type DeleteABTestSceneResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestSceneResponse) SetHeaders(v map[string]*string) *DeleteABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestSceneResponse) SetBody(v *DeleteABTestSceneResponseBody) *DeleteABTestSceneResponse {
	s.Body = v
	return s
}

type DeleteFunctionInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64  `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceResponseBody) SetCode(v string) *DeleteFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetHttpCode(v int64) *DeleteFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetLatency(v int64) *DeleteFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetMessage(v string) *DeleteFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetRequestId(v string) *DeleteFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetStatus(v string) *DeleteFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type DeleteFunctionInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceResponse) SetHeaders(v map[string]*string) *DeleteFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionInstanceResponse) SetBody(v *DeleteFunctionInstanceResponseBody) *DeleteFunctionInstanceResponse {
	s.Body = v
	return s
}

type DeleteModelResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) SetResult(v string) *DeleteModelResponseBody {
	s.Result = &v
	return s
}

type DeleteModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

type DeleteSortScriptResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptResponseBody) SetRequestId(v string) *DeleteSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSortScriptResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptResponse) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptResponse) SetHeaders(v map[string]*string) *DeleteSortScriptResponse {
	s.Headers = v
	return s
}

func (s *DeleteSortScriptResponse) SetBody(v *DeleteSortScriptResponseBody) *DeleteSortScriptResponse {
	s.Body = v
	return s
}

type DeleteSortScriptFileResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSortScriptFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptFileResponseBody) SetRequestId(v string) *DeleteSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSortScriptFileResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSortScriptFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptFileResponse) SetHeaders(v map[string]*string) *DeleteSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSortScriptFileResponse) SetBody(v *DeleteSortScriptFileResponseBody) *DeleteSortScriptFileResponse {
	s.Body = v
	return s
}

type DescribeABTestExperimentResponseBody struct {
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBody) SetRequestId(v string) *DescribeABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestExperimentResponseBody) SetResult(v *DescribeABTestExperimentResponseBodyResult) *DescribeABTestExperimentResponseBody {
	s.Result = v
	return s
}

type DescribeABTestExperimentResponseBodyResult struct {
	Created *int32                                            `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string                                           `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	Online  *bool                                             `json:"online,omitempty" xml:"online,omitempty"`
	Params  *DescribeABTestExperimentResponseBodyResultParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	Traffic *int32                                            `json:"traffic,omitempty" xml:"traffic,omitempty"`
	Updated *int32                                            `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeABTestExperimentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBodyResult) SetCreated(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetId(v string) *DescribeABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetName(v string) *DescribeABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetOnline(v bool) *DescribeABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetParams(v *DescribeABTestExperimentResponseBodyResultParams) *DescribeABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetTraffic(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetUpdated(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeABTestExperimentResponseBodyResultParams struct {
	FirstFormulaName *string `json:"first_formula_name,omitempty" xml:"first_formula_name,omitempty"`
}

func (s DescribeABTestExperimentResponseBodyResultParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponseBodyResultParams) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBodyResultParams) SetFirstFormulaName(v string) *DescribeABTestExperimentResponseBodyResultParams {
	s.FirstFormulaName = &v
	return s
}

type DescribeABTestExperimentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponse) SetHeaders(v map[string]*string) *DescribeABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestExperimentResponse) SetBody(v *DescribeABTestExperimentResponseBody) *DescribeABTestExperimentResponse {
	s.Body = v
	return s
}

type DescribeABTestGroupResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponseBody) SetRequestId(v string) *DescribeABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestGroupResponseBody) SetResult(v *DescribeABTestGroupResponseBodyResult) *DescribeABTestGroupResponseBody {
	s.Result = v
	return s
}

type DescribeABTestGroupResponseBodyResult struct {
	Created *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeABTestGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponseBodyResult) SetCreated(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetId(v string) *DescribeABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetName(v string) *DescribeABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetStatus(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetUpdated(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeABTestGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponse) SetHeaders(v map[string]*string) *DescribeABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestGroupResponse) SetBody(v *DescribeABTestGroupResponseBody) *DescribeABTestGroupResponse {
	s.Body = v
	return s
}

type DescribeABTestSceneResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponseBody) SetRequestId(v string) *DescribeABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestSceneResponseBody) SetResult(v *DescribeABTestSceneResponseBodyResult) *DescribeABTestSceneResponseBody {
	s.Result = v
	return s
}

type DescribeABTestSceneResponseBodyResult struct {
	Created *int32    `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string   `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string   `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32    `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32    `json:"updated,omitempty" xml:"updated,omitempty"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s DescribeABTestSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponseBodyResult) SetCreated(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetId(v string) *DescribeABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetName(v string) *DescribeABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetStatus(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetUpdated(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetValues(v []*string) *DescribeABTestSceneResponseBodyResult {
	s.Values = v
	return s
}

type DescribeABTestSceneResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponse) SetHeaders(v map[string]*string) *DescribeABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestSceneResponse) SetBody(v *DescribeABTestSceneResponseBody) *DescribeABTestSceneResponse {
	s.Body = v
	return s
}

type DescribeAppResponseBody struct {
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBody) SetRequestId(v string) *DescribeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppResponseBody) SetResult(v *DescribeAppResponseBodyResult) *DescribeAppResponseBody {
	s.Result = v
	return s
}

type DescribeAppResponseBodyResult struct {
	AlgoDeploymentId *int32                               `json:"algoDeploymentId,omitempty" xml:"algoDeploymentId,omitempty"`
	AutoSwitch       *bool                                `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	ClusterName      *string                              `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	Created          *int32                               `json:"created,omitempty" xml:"created,omitempty"`
	Description      *string                              `json:"description,omitempty" xml:"description,omitempty"`
	Domain           *DescribeAppResponseBodyResultDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	FetchFields      []*string                            `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	Id               *string                              `json:"id,omitempty" xml:"id,omitempty"`
	ProgressPercent  *int32                               `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	Quota            *DescribeAppResponseBodyResultQuota  `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	Schema           map[string]interface{}               `json:"schema,omitempty" xml:"schema,omitempty"`
	Status           *string                              `json:"status,omitempty" xml:"status,omitempty"`
	Type             *string                              `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResult) SetAlgoDeploymentId(v int32) *DescribeAppResponseBodyResult {
	s.AlgoDeploymentId = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetAutoSwitch(v bool) *DescribeAppResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetClusterName(v string) *DescribeAppResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetCreated(v int32) *DescribeAppResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDescription(v string) *DescribeAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDomain(v *DescribeAppResponseBodyResultDomain) *DescribeAppResponseBodyResult {
	s.Domain = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetFetchFields(v []*string) *DescribeAppResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetId(v string) *DescribeAppResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetProgressPercent(v int32) *DescribeAppResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetQuota(v *DescribeAppResponseBodyResultQuota) *DescribeAppResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSchema(v map[string]interface{}) *DescribeAppResponseBodyResult {
	s.Schema = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetStatus(v string) *DescribeAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetType(v string) *DescribeAppResponseBodyResult {
	s.Type = &v
	return s
}

type DescribeAppResponseBodyResultDomain struct {
	Category  *string                                       `json:"category,omitempty" xml:"category,omitempty"`
	Functions *DescribeAppResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	Name      *string                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppResponseBodyResultDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDomain) SetCategory(v string) *DescribeAppResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) SetFunctions(v *DescribeAppResponseBodyResultDomainFunctions) *DescribeAppResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) SetName(v string) *DescribeAppResponseBodyResultDomain {
	s.Name = &v
	return s
}

type DescribeAppResponseBodyResultDomainFunctions struct {
	Algo    []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	Qp      []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s DescribeAppResponseBodyResultDomainFunctions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetAlgo(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetQp(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetService(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

type DescribeAppResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Qps             *int32  `json:"qps,omitempty" xml:"qps,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeAppResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetQps(v int32) *DescribeAppResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetSpec(v string) *DescribeAppResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type DescribeAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppResponse) SetHeaders(v map[string]*string) *DescribeAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppResponse) SetBody(v *DescribeAppResponseBody) *DescribeAppResponse {
	s.Body = v
	return s
}

type DescribeAppGroupResponseBody struct {
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBody) SetRequestId(v string) *DescribeAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppGroupResponseBody) SetResult(v *DescribeAppGroupResponseBodyResult) *DescribeAppGroupResponseBody {
	s.Result = v
	return s
}

type DescribeAppGroupResponseBodyResult struct {
	ChargeType                        *string                                  `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay                       *int32                                   `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	CommodityCode                     *string                                  `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Created                           *int32                                   `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion                    *string                                  `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Description                       *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	Domain                            *string                                  `json:"domain,omitempty" xml:"domain,omitempty"`
	ExpireOn                          *string                                  `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	FirstRankAlgoDeploymentId         *int32                                   `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	HasPendingQuotaReviewTask         *int32                                   `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	Id                                *string                                  `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId                        *string                                  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockMode                          *string                                  `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration                *int32                                   `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	Name                              *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	PendingSecondRankAlgoDeploymentId *int32                                   `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	ProcessingOrderId                 *string                                  `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	Produced                          *int32                                   `json:"produced,omitempty" xml:"produced,omitempty"`
	ProjectId                         *string                                  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Quota                             *DescribeAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	ResourceGroupId                   *string                                  `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	SecondRankAlgoDeploymentId        *int32                                   `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	Status                            *string                                  `json:"status,omitempty" xml:"status,omitempty"`
	SwitchedTime                      *int32                                   `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Type                              *string                                  `json:"type,omitempty" xml:"type,omitempty"`
	Updated                           *int32                                   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResult) SetChargeType(v string) *DescribeAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetChargingWay(v int32) *DescribeAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCommodityCode(v string) *DescribeAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCreated(v int32) *DescribeAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCurrentVersion(v string) *DescribeAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetDescription(v string) *DescribeAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetDomain(v string) *DescribeAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetExpireOn(v string) *DescribeAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *DescribeAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetId(v string) *DescribeAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetInstanceId(v string) *DescribeAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetLockMode(v string) *DescribeAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *DescribeAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetName(v string) *DescribeAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProcessingOrderId(v string) *DescribeAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProduced(v int32) *DescribeAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProjectId(v string) *DescribeAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetQuota(v *DescribeAppGroupResponseBodyResultQuota) *DescribeAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetResourceGroupId(v string) *DescribeAppGroupResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetStatus(v string) *DescribeAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetSwitchedTime(v int32) *DescribeAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetType(v string) *DescribeAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetUpdated(v int32) *DescribeAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeAppGroupResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeAppGroupResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetSpec(v string) *DescribeAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type DescribeAppGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponse) SetHeaders(v map[string]*string) *DescribeAppGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppGroupResponse) SetBody(v *DescribeAppGroupResponseBody) *DescribeAppGroupResponse {
	s.Body = v
	return s
}

type DescribeAppGroupDataReportRequest struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s DescribeAppGroupDataReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDataReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDataReportRequest) SetEndTime(v string) *DescribeAppGroupDataReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAppGroupDataReportRequest) SetStartTime(v string) *DescribeAppGroupDataReportRequest {
	s.StartTime = &v
	return s
}

type DescribeAppGroupDataReportResponseBody struct {
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeAppGroupDataReportResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppGroupDataReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDataReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDataReportResponseBody) SetRequestId(v string) *DescribeAppGroupDataReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBody) SetResult(v *DescribeAppGroupDataReportResponseBodyResult) *DescribeAppGroupDataReportResponseBody {
	s.Result = v
	return s
}

type DescribeAppGroupDataReportResponseBodyResult struct {
	ReceivedCount  *int32                                                        `json:"receivedCount,omitempty" xml:"receivedCount,omitempty"`
	ReceivedSample []*DescribeAppGroupDataReportResponseBodyResultReceivedSample `json:"receivedSample,omitempty" xml:"receivedSample,omitempty" type:"Repeated"`
}

func (s DescribeAppGroupDataReportResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDataReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDataReportResponseBodyResult) SetReceivedCount(v int32) *DescribeAppGroupDataReportResponseBodyResult {
	s.ReceivedCount = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResult) SetReceivedSample(v []*DescribeAppGroupDataReportResponseBodyResultReceivedSample) *DescribeAppGroupDataReportResponseBodyResult {
	s.ReceivedSample = v
	return s
}

type DescribeAppGroupDataReportResponseBodyResultReceivedSample struct {
	Message        *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
	ReceivedTimeMs *int64                                                             `json:"receivedTimeMs,omitempty" xml:"receivedTimeMs,omitempty"`
}

func (s DescribeAppGroupDataReportResponseBodyResultReceivedSample) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDataReportResponseBodyResultReceivedSample) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSample) SetMessage(v *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) *DescribeAppGroupDataReportResponseBodyResultReceivedSample {
	s.Message = v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSample) SetReceivedTimeMs(v int64) *DescribeAppGroupDataReportResponseBodyResultReceivedSample {
	s.ReceivedTimeMs = &v
	return s
}

type DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage struct {
	Arg1       *string `json:"arg1,omitempty" xml:"arg1,omitempty"`
	Arg3       *string `json:"arg3,omitempty" xml:"arg3,omitempty"`
	Args       *string `json:"args,omitempty" xml:"args,omitempty"`
	ClientIp   *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
	EventId    *int32  `json:"eventId,omitempty" xml:"eventId,omitempty"`
	Page       *string `json:"page,omitempty" xml:"page,omitempty"`
	SdkType    *string `json:"sdkType,omitempty" xml:"sdkType,omitempty"`
	SdkVersion *string `json:"sdkVersion,omitempty" xml:"sdkVersion,omitempty"`
	SessionId  *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetArg1(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.Arg1 = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetArg3(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.Arg3 = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetArgs(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.Args = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetClientIp(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.ClientIp = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetEventId(v int32) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.EventId = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetPage(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.Page = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetSdkType(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.SdkType = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetSdkVersion(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.SdkVersion = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetSessionId(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.SessionId = &v
	return s
}

func (s *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage) SetUserId(v string) *DescribeAppGroupDataReportResponseBodyResultReceivedSampleMessage {
	s.UserId = &v
	return s
}

type DescribeAppGroupDataReportResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppGroupDataReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppGroupDataReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDataReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDataReportResponse) SetHeaders(v map[string]*string) *DescribeAppGroupDataReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppGroupDataReportResponse) SetBody(v *DescribeAppGroupDataReportResponseBody) *DescribeAppGroupDataReportResponse {
	s.Body = v
	return s
}

type DescribeAppGroupStatisticsResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeAppGroupStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupStatisticsResponseBody) SetRequestId(v string) *DescribeAppGroupStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppGroupStatisticsResponseBody) SetResult(v map[string]interface{}) *DescribeAppGroupStatisticsResponseBody {
	s.Result = v
	return s
}

type DescribeAppGroupStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppGroupStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppGroupStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAppGroupStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppGroupStatisticsResponse) SetBody(v *DescribeAppGroupStatisticsResponseBody) *DescribeAppGroupStatisticsResponse {
	s.Body = v
	return s
}

type DescribeAppStatisticsResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeAppStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppStatisticsResponseBody) SetRequestId(v string) *DescribeAppStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppStatisticsResponseBody) SetResult(v map[string]interface{}) *DescribeAppStatisticsResponseBody {
	s.Result = v
	return s
}

type DescribeAppStatisticsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAppStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppStatisticsResponse) SetBody(v *DescribeAppStatisticsResponseBody) *DescribeAppStatisticsResponse {
	s.Body = v
	return s
}

type DescribeAppsResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetResult(v []map[string]interface{}) *DescribeAppsResponseBody {
	s.Result = v
	return s
}

type DescribeAppsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeDataCollctionResponseBody struct {
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeDataCollctionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeDataCollctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCollctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponseBody) SetRequestId(v string) *DescribeDataCollctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCollctionResponseBody) SetResult(v *DescribeDataCollctionResponseBodyResult) *DescribeDataCollctionResponseBody {
	s.Result = v
	return s
}

type DescribeDataCollctionResponseBodyResult struct {
	Created            *int32  `json:"created,omitempty" xml:"created,omitempty"`
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	Id                 *string `json:"id,omitempty" xml:"id,omitempty"`
	IndustryName       *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Status             *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SundialId          *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated            *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeDataCollctionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCollctionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponseBodyResult) SetCreated(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetDataCollectionType(v string) *DescribeDataCollctionResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetId(v string) *DescribeDataCollctionResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetIndustryName(v string) *DescribeDataCollctionResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetName(v string) *DescribeDataCollctionResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetStatus(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetSundialId(v string) *DescribeDataCollctionResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetType(v string) *DescribeDataCollctionResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetUpdated(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeDataCollctionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataCollctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataCollctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCollctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponse) SetHeaders(v map[string]*string) *DescribeDataCollctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCollctionResponse) SetBody(v *DescribeDataCollctionResponseBody) *DescribeDataCollctionResponse {
	s.Body = v
	return s
}

type DescribeFirstRankResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBody) SetRequestId(v string) *DescribeFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirstRankResponseBody) SetResult(v *DescribeFirstRankResponseBodyResult) *DescribeFirstRankResponseBody {
	s.Result = v
	return s
}

type DescribeFirstRankResponseBodyResult struct {
	Active      *bool                                      `json:"active,omitempty" xml:"active,omitempty"`
	Description *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	Meta        []*DescribeFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name        *string                                    `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBodyResult) SetActive(v bool) *DescribeFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetDescription(v string) *DescribeFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetMeta(v []*DescribeFirstRankResponseBodyResultMeta) *DescribeFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetName(v string) *DescribeFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type DescribeFirstRankResponseBodyResultMeta struct {
	Arg       *string  `json:"arg,omitempty" xml:"arg,omitempty"`
	Attribute *string  `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Weight    *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DescribeFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetArg(v string) *DescribeFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetAttribute(v string) *DescribeFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetWeight(v float32) *DescribeFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type DescribeFirstRankResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponse) SetHeaders(v map[string]*string) *DescribeFirstRankResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirstRankResponse) SetBody(v *DescribeFirstRankResponseBody) *DescribeFirstRankResponse {
	s.Body = v
	return s
}

type DescribeInterventionDictionaryResponseBody struct {
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeInterventionDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponseBody) SetRequestId(v string) *DescribeInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBody) SetResult(v *DescribeInterventionDictionaryResponseBodyResult) *DescribeInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

type DescribeInterventionDictionaryResponseBodyResult struct {
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	Created  *string `json:"created,omitempty" xml:"created,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated  *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeInterventionDictionaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetCreated(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetName(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetType(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetUpdated(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeInterventionDictionaryResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInterventionDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponse) SetHeaders(v map[string]*string) *DescribeInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterventionDictionaryResponse) SetBody(v *DescribeInterventionDictionaryResponseBody) *DescribeInterventionDictionaryResponse {
	s.Body = v
	return s
}

type DescribeModelResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelResponseBody) SetRequestId(v string) *DescribeModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelResponseBody) SetResult(v map[string]interface{}) *DescribeModelResponseBody {
	s.Result = v
	return s
}

type DescribeModelResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelResponse) SetHeaders(v map[string]*string) *DescribeModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelResponse) SetBody(v *DescribeModelResponseBody) *DescribeModelResponse {
	s.Body = v
	return s
}

type DescribeQueryProcessorResponseBody struct {
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponseBody) SetRequestId(v string) *DescribeQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryProcessorResponseBody) SetResult(v *DescribeQueryProcessorResponseBodyResult) *DescribeQueryProcessorResponseBody {
	s.Result = v
	return s
}

type DescribeQueryProcessorResponseBodyResult struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Created    *int32                   `json:"created,omitempty" xml:"created,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	Updated    *int32                   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeQueryProcessorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponseBodyResult) SetActive(v bool) *DescribeQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetCreated(v int32) *DescribeQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetDomain(v string) *DescribeQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetIndexes(v []*string) *DescribeQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetName(v string) *DescribeQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *DescribeQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetUpdated(v int32) *DescribeQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeQueryProcessorResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponse) SetHeaders(v map[string]*string) *DescribeQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryProcessorResponse) SetBody(v *DescribeQueryProcessorResponseBody) *DescribeQueryProcessorResponse {
	s.Body = v
	return s
}

type DescribeRegionResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeRegionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionResponseBody) SetRequestId(v string) *DescribeRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionResponseBody) SetResult(v *DescribeRegionResponseBodyResult) *DescribeRegionResponseBody {
	s.Result = v
	return s
}

type DescribeRegionResponseBodyResult struct {
	Config   map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	RegionId *string                `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionResponseBodyResult) SetConfig(v map[string]interface{}) *DescribeRegionResponseBodyResult {
	s.Config = v
	return s
}

func (s *DescribeRegionResponseBodyResult) SetRegionId(v string) *DescribeRegionResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescribeRegionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionResponse) SetHeaders(v map[string]*string) *DescribeRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionResponse) SetBody(v *DescribeRegionResponseBody) *DescribeRegionResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*DescribeRegionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody {
	s.Result = v
	return s
}

type DescribeRegionsResponseBodyResult struct {
	ConsoleUrl *string `json:"consoleUrl,omitempty" xml:"consoleUrl,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LocalName  *string `json:"localName,omitempty" xml:"localName,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) SetConsoleUrl(v string) *DescribeRegionsResponseBodyResult {
	s.ConsoleUrl = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeScheduledTaskResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskResponseBody) SetRequestId(v string) *DescribeScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTaskResponseBody) SetResult(v map[string]interface{}) *DescribeScheduledTaskResponseBody {
	s.Result = v
	return s
}

type DescribeScheduledTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskResponse) SetHeaders(v map[string]*string) *DescribeScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTaskResponse) SetBody(v *DescribeScheduledTaskResponseBody) *DescribeScheduledTaskResponse {
	s.Body = v
	return s
}

type DescribeSecondRankResponseBody struct {
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeSecondRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeSecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponseBody) SetRequestId(v string) *DescribeSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecondRankResponseBody) SetResult(v *DescribeSecondRankResponseBodyResult) *DescribeSecondRankResponseBody {
	s.Result = v
	return s
}

type DescribeSecondRankResponseBodyResult struct {
	Active      *bool   `json:"active,omitempty" xml:"active,omitempty"`
	Created     *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDefault   *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	IsSys       *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	Meta        *string `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated     *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeSecondRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecondRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponseBodyResult) SetActive(v bool) *DescribeSecondRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetCreated(v int32) *DescribeSecondRankResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetDescription(v string) *DescribeSecondRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetId(v string) *DescribeSecondRankResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetIsDefault(v string) *DescribeSecondRankResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetIsSys(v string) *DescribeSecondRankResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetMeta(v string) *DescribeSecondRankResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetName(v string) *DescribeSecondRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetUpdated(v int32) *DescribeSecondRankResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeSecondRankResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecondRankResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponse) SetHeaders(v map[string]*string) *DescribeSecondRankResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecondRankResponse) SetBody(v *DescribeSecondRankResponseBody) *DescribeSecondRankResponse {
	s.Body = v
	return s
}

type DescribeSlowQueryStatusResponseBody struct {
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeSlowQueryStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeSlowQueryStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowQueryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponseBody) SetRequestId(v string) *DescribeSlowQueryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBody) SetResult(v *DescribeSlowQueryStatusResponseBodyResult) *DescribeSlowQueryStatusResponseBody {
	s.Result = v
	return s
}

type DescribeSlowQueryStatusResponseBodyResult struct {
	AppGroupId *string `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeSlowQueryStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowQueryStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetAppGroupId(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetRegion(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetStatus(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.Status = &v
	return s
}

type DescribeSlowQueryStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowQueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowQueryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponse) SetHeaders(v map[string]*string) *DescribeSlowQueryStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowQueryStatusResponse) SetBody(v *DescribeSlowQueryStatusResponseBody) *DescribeSlowQueryStatusResponse {
	s.Body = v
	return s
}

type DescribeUserAnalyzerRequest struct {
	With *string `json:"with,omitempty" xml:"with,omitempty"`
}

func (s DescribeUserAnalyzerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerRequest) SetWith(v string) *DescribeUserAnalyzerRequest {
	s.With = &v
	return s
}

type DescribeUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerResponseBody) SetRequestId(v string) *DescribeUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *DescribeUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type DescribeUserAnalyzerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerResponse) SetHeaders(v map[string]*string) *DescribeUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAnalyzerResponse) SetBody(v *DescribeUserAnalyzerResponseBody) *DescribeUserAnalyzerResponse {
	s.Body = v
	return s
}

type DisableSlowQueryResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DisableSlowQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSlowQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSlowQueryResponseBody) SetRequestId(v string) *DisableSlowQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSlowQueryResponseBody) SetResult(v map[string]interface{}) *DisableSlowQueryResponseBody {
	s.Result = v
	return s
}

type DisableSlowQueryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSlowQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSlowQueryResponse) GoString() string {
	return s.String()
}

func (s *DisableSlowQueryResponse) SetHeaders(v map[string]*string) *DisableSlowQueryResponse {
	s.Headers = v
	return s
}

func (s *DisableSlowQueryResponse) SetBody(v *DisableSlowQueryResponseBody) *DisableSlowQueryResponse {
	s.Body = v
	return s
}

type EnableSlowQueryResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s EnableSlowQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSlowQueryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSlowQueryResponseBody) SetRequestId(v string) *EnableSlowQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSlowQueryResponseBody) SetResult(v map[string]interface{}) *EnableSlowQueryResponseBody {
	s.Result = v
	return s
}

type EnableSlowQueryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSlowQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSlowQueryResponse) GoString() string {
	return s.String()
}

func (s *EnableSlowQueryResponse) SetHeaders(v map[string]*string) *EnableSlowQueryResponse {
	s.Headers = v
	return s
}

func (s *EnableSlowQueryResponse) SetBody(v *EnableSlowQueryResponseBody) *EnableSlowQueryResponse {
	s.Body = v
	return s
}

type GenerateMergedTableRequest struct {
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s GenerateMergedTableRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableRequest) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableRequest) SetSpec(v string) *GenerateMergedTableRequest {
	s.Spec = &v
	return s
}

type GenerateMergedTableResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GenerateMergedTableResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GenerateMergedTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponseBody) SetRequestId(v string) *GenerateMergedTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMergedTableResponseBody) SetResult(v *GenerateMergedTableResponseBodyResult) *GenerateMergedTableResponseBody {
	s.Result = v
	return s
}

type GenerateMergedTableResponseBodyResult struct {
	FromTable  map[string]interface{} `json:"fromTable,omitempty" xml:"fromTable,omitempty"`
	MergeTable map[string]interface{} `json:"mergeTable,omitempty" xml:"mergeTable,omitempty"`
	PrimaryKey *string                `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
}

func (s GenerateMergedTableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponseBodyResult) SetFromTable(v map[string]interface{}) *GenerateMergedTableResponseBodyResult {
	s.FromTable = v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) SetMergeTable(v map[string]interface{}) *GenerateMergedTableResponseBodyResult {
	s.MergeTable = v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) SetPrimaryKey(v string) *GenerateMergedTableResponseBodyResult {
	s.PrimaryKey = &v
	return s
}

type GenerateMergedTableResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateMergedTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateMergedTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableResponse) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponse) SetHeaders(v map[string]*string) *GenerateMergedTableResponse {
	s.Headers = v
	return s
}

func (s *GenerateMergedTableResponse) SetBody(v *GenerateMergedTableResponseBody) *GenerateMergedTableResponse {
	s.Body = v
	return s
}

type GetDomainRequest struct {
	AppGroupIdentity *string `json:"appGroupIdentity,omitempty" xml:"appGroupIdentity,omitempty"`
}

func (s GetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) SetAppGroupIdentity(v string) *GetDomainRequest {
	s.AppGroupIdentity = &v
	return s
}

type GetDomainResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainResponseBody) SetResult(v map[string]interface{}) *GetDomainResponseBody {
	s.Result = v
	return s
}

type GetDomainResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetBody(v *GetDomainResponseBody) *GetDomainResponse {
	s.Body = v
	return s
}

type GetFunctionCurrentVersionRequest struct {
	Category     *string `json:"category,omitempty" xml:"category,omitempty"`
	Domain       *string `json:"domain,omitempty" xml:"domain,omitempty"`
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	ModelType    *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
}

func (s GetFunctionCurrentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionRequest) SetCategory(v string) *GetFunctionCurrentVersionRequest {
	s.Category = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetDomain(v string) *GetFunctionCurrentVersionRequest {
	s.Domain = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetFunctionType(v string) *GetFunctionCurrentVersionRequest {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetModelType(v string) *GetFunctionCurrentVersionRequest {
	s.ModelType = &v
	return s
}

type GetFunctionCurrentVersionResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64                                       `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64                                       `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetFunctionCurrentVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status    *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBody) SetCode(v string) *GetFunctionCurrentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetHttpCode(v int64) *GetFunctionCurrentVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetLatency(v int64) *GetFunctionCurrentVersionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetMessage(v string) *GetFunctionCurrentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetRequestId(v string) *GetFunctionCurrentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetResult(v *GetFunctionCurrentVersionResponseBodyResult) *GetFunctionCurrentVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetStatus(v string) *GetFunctionCurrentVersionResponseBody {
	s.Status = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResult struct {
	FunctionName  *string                                                   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionType  *string                                                   `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	ModelType     *string                                                   `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	VersionConfig *GetFunctionCurrentVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	VersionId     *int64                                                    `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionName   *string                                                   `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetFunctionName(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetFunctionType(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetModelType(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionConfig(v *GetFunctionCurrentVersionResponseBodyResultVersionConfig) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionConfig = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionId(v int64) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionName(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionName = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfig struct {
	CreateParameters []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	Depends          []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends          `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	UsageParameters  []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters  `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfig) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetCreateParameters(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetDepends(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.Depends = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetUsageParameters(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.UsageParameters = v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) SetName(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) SetRequired(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	s.Required = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends struct {
	Condition   *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Dependency  *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetCondition(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Condition = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetDependency(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Dependency = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetDescription(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Description = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) SetName(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) SetRequired(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	s.Required = &v
	return s
}

type GetFunctionCurrentVersionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionCurrentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionCurrentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponse) SetHeaders(v map[string]*string) *GetFunctionCurrentVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionCurrentVersionResponse) SetBody(v *GetFunctionCurrentVersionResponseBody) *GetFunctionCurrentVersionResponse {
	s.Body = v
	return s
}

type GetFunctionDefaultInstanceResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 功能名称
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	HttpCode     *int64  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 实例名称
	InstanceName *string                                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Latency      *int64                                        `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetFunctionDefaultInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status       *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionDefaultInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponseBody) SetCode(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetFunctionName(v string) *GetFunctionDefaultInstanceResponseBody {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetHttpCode(v int64) *GetFunctionDefaultInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetLatency(v int64) *GetFunctionDefaultInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetMessage(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetRequestId(v string) *GetFunctionDefaultInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetResult(v *GetFunctionDefaultInstanceResponseBodyResult) *GetFunctionDefaultInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetStatus(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Status = &v
	return s
}

type GetFunctionDefaultInstanceResponseBodyResult struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetFunctionDefaultInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

type GetFunctionDefaultInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionDefaultInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponse) SetHeaders(v map[string]*string) *GetFunctionDefaultInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionDefaultInstanceResponse) SetBody(v *GetFunctionDefaultInstanceResponseBody) *GetFunctionDefaultInstanceResponse {
	s.Body = v
	return s
}

type GetFunctionInstanceRequest struct {
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
}

func (s GetFunctionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceRequest) SetOutput(v string) *GetFunctionInstanceRequest {
	s.Output = &v
	return s
}

type GetFunctionInstanceResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64                                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64                                 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetFunctionInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status    *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBody) SetCode(v string) *GetFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetHttpCode(v int64) *GetFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetLatency(v int64) *GetFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetMessage(v string) *GetFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetRequestId(v string) *GetFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetResult(v *GetFunctionInstanceResponseBodyResult) *GetFunctionInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetStatus(v string) *GetFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type GetFunctionInstanceResponseBodyResult struct {
	Belongs          *GetFunctionInstanceResponseBodyResultBelongs            `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	CreateParameters []*GetFunctionInstanceResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	CreateTime       *int64                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Cron             *string                                                  `json:"Cron,omitempty" xml:"Cron,omitempty"`
	Description      *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ExtendInfo       *string                                                  `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FunctionName     *string                                                  `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionType     *string                                                  `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	InstanceName     *string                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ModelType        *string                                                  `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	Source           *string                                                  `json:"Source,omitempty" xml:"Source,omitempty"`
	Status           *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Task             *GetFunctionInstanceResponseBodyResultTask               `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	VersionId        *int64                                                   `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResult) SetBelongs(v *GetFunctionInstanceResponseBodyResultBelongs) *GetFunctionInstanceResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateParameters(v []*GetFunctionInstanceResponseBodyResultCreateParameters) *GetFunctionInstanceResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateTime(v int64) *GetFunctionInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCron(v string) *GetFunctionInstanceResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetDescription(v string) *GetFunctionInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetExtendInfo(v string) *GetFunctionInstanceResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionName(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionType(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetModelType(v string) *GetFunctionInstanceResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetSource(v string) *GetFunctionInstanceResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetStatus(v string) *GetFunctionInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetTask(v *GetFunctionInstanceResponseBodyResultTask) *GetFunctionInstanceResponseBodyResult {
	s.Task = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetVersionId(v int64) *GetFunctionInstanceResponseBodyResult {
	s.VersionId = &v
	return s
}

type GetFunctionInstanceResponseBodyResultBelongs struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultBelongs) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetCategory(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetDomain(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetLanguage(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Language = &v
	return s
}

type GetFunctionInstanceResponseBodyResultCreateParameters struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetName(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetValue(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

type GetFunctionInstanceResponseBodyResultTask struct {
	DagStatus   *string `json:"DagStatus,omitempty" xml:"DagStatus,omitempty"`
	LastRunTime *int64  `json:"LastRunTime,omitempty" xml:"LastRunTime,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultTask) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultTask) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetDagStatus(v string) *GetFunctionInstanceResponseBodyResultTask {
	s.DagStatus = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetLastRunTime(v int64) *GetFunctionInstanceResponseBodyResultTask {
	s.LastRunTime = &v
	return s
}

type GetFunctionInstanceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponse) SetHeaders(v map[string]*string) *GetFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionInstanceResponse) SetBody(v *GetFunctionInstanceResponseBody) *GetFunctionInstanceResponse {
	s.Body = v
	return s
}

type GetFunctionVersionResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64                                `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetFunctionVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status    *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBody) SetCode(v string) *GetFunctionVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetHttpCode(v int64) *GetFunctionVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetLatency(v int64) *GetFunctionVersionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetMessage(v string) *GetFunctionVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetRequestId(v string) *GetFunctionVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetResult(v *GetFunctionVersionResponseBodyResult) *GetFunctionVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionVersionResponseBody) SetStatus(v string) *GetFunctionVersionResponseBody {
	s.Status = &v
	return s
}

type GetFunctionVersionResponseBodyResult struct {
	FunctionName  *string                                            `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionType  *string                                            `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	ModelType     *string                                            `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	VersionConfig *GetFunctionVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	VersionId     *int64                                             `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionName   *string                                            `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetFunctionVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResult) SetFunctionName(v string) *GetFunctionVersionResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetFunctionType(v string) *GetFunctionVersionResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetModelType(v string) *GetFunctionVersionResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionConfig(v *GetFunctionVersionResponseBodyResultVersionConfig) *GetFunctionVersionResponseBodyResult {
	s.VersionConfig = v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionId(v int64) *GetFunctionVersionResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionName(v string) *GetFunctionVersionResponseBodyResult {
	s.VersionName = &v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfig struct {
	CreateParameters []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	Depends          []*GetFunctionVersionResponseBodyResultVersionConfigDepends          `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	UsageParameters  []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters  `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfig) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetCreateParameters(v []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetDepends(v []*GetFunctionVersionResponseBodyResultVersionConfigDepends) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.Depends = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetUsageParameters(v []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.UsageParameters = v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfigCreateParameters struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) SetName(v string) *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) SetRequired(v string) *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	s.Required = &v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfigDepends struct {
	Condition   *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Dependency  *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigDepends) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigDepends) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetCondition(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Condition = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetDependency(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Dependency = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetDescription(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Description = &v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfigUsageParameters struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) SetName(v string) *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) SetRequired(v string) *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	s.Required = &v
	return s
}

type GetFunctionVersionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponse) SetHeaders(v map[string]*string) *GetFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionVersionResponse) SetBody(v *GetFunctionVersionResponseBody) *GetFunctionVersionResponse {
	s.Body = v
	return s
}

type GetModelProgressResponseBody struct {
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetModelProgressResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetModelProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelProgressResponseBody) SetRequestId(v string) *GetModelProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelProgressResponseBody) SetResult(v *GetModelProgressResponseBodyResult) *GetModelProgressResponseBody {
	s.Result = v
	return s
}

type GetModelProgressResponseBodyResult struct {
	Progress *int32  `json:"progress,omitempty" xml:"progress,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetModelProgressResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetModelProgressResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetModelProgressResponseBodyResult) SetProgress(v int32) *GetModelProgressResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetModelProgressResponseBodyResult) SetStatus(v string) *GetModelProgressResponseBodyResult {
	s.Status = &v
	return s
}

type GetModelProgressResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelProgressResponse) GoString() string {
	return s.String()
}

func (s *GetModelProgressResponse) SetHeaders(v map[string]*string) *GetModelProgressResponse {
	s.Headers = v
	return s
}

func (s *GetModelProgressResponse) SetBody(v *GetModelProgressResponseBody) *GetModelProgressResponse {
	s.Body = v
	return s
}

type GetModelReportResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetModelReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelReportResponseBody) SetRequestId(v string) *GetModelReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelReportResponseBody) SetResult(v map[string]interface{}) *GetModelReportResponseBody {
	s.Result = v
	return s
}

type GetModelReportResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelReportResponse) GoString() string {
	return s.String()
}

func (s *GetModelReportResponse) SetHeaders(v map[string]*string) *GetModelReportResponse {
	s.Headers = v
	return s
}

func (s *GetModelReportResponse) SetBody(v *GetModelReportResponseBody) *GetModelReportResponse {
	s.Body = v
	return s
}

type GetScriptFileNamesResponseBody struct {
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetScriptFileNamesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetScriptFileNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScriptFileNamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponseBody) SetRequestId(v string) *GetScriptFileNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptFileNamesResponseBody) SetResult(v []*GetScriptFileNamesResponseBodyResult) *GetScriptFileNamesResponseBody {
	s.Result = v
	return s
}

type GetScriptFileNamesResponseBodyResult struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FileName   *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s GetScriptFileNamesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetScriptFileNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponseBodyResult) SetCreateTime(v string) *GetScriptFileNamesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetFileName(v string) *GetScriptFileNamesResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetModifyTime(v string) *GetScriptFileNamesResponseBodyResult {
	s.ModifyTime = &v
	return s
}

type GetScriptFileNamesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScriptFileNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScriptFileNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScriptFileNamesResponse) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponse) SetHeaders(v map[string]*string) *GetScriptFileNamesResponse {
	s.Headers = v
	return s
}

func (s *GetScriptFileNamesResponse) SetBody(v *GetScriptFileNamesResponseBody) *GetScriptFileNamesResponse {
	s.Body = v
	return s
}

type GetSearchStrategyResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchStrategyResponseBody) SetRequestId(v string) *GetSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type GetSearchStrategyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetSearchStrategyResponse) SetHeaders(v map[string]*string) *GetSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetSearchStrategyResponse) SetBody(v *GetSearchStrategyResponseBody) *GetSearchStrategyResponse {
	s.Body = v
	return s
}

type GetSortScriptResponseBody struct {
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSortScriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponseBody) SetRequestId(v string) *GetSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSortScriptResponseBody) SetResult(v *GetSortScriptResponseBodyResult) *GetSortScriptResponseBody {
	s.Result = v
	return s
}

type GetSortScriptResponseBodyResult struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Scope      *string `json:"scope,omitempty" xml:"scope,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSortScriptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponseBodyResult) SetCreateTime(v string) *GetSortScriptResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetModifyTime(v string) *GetSortScriptResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetScope(v string) *GetSortScriptResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetStatus(v string) *GetSortScriptResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetType(v string) *GetSortScriptResponseBodyResult {
	s.Type = &v
	return s
}

type GetSortScriptResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptResponse) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponse) SetHeaders(v map[string]*string) *GetSortScriptResponse {
	s.Headers = v
	return s
}

func (s *GetSortScriptResponse) SetBody(v *GetSortScriptResponseBody) *GetSortScriptResponse {
	s.Body = v
	return s
}

type GetSortScriptFileResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSortScriptFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSortScriptFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponseBody) SetRequestId(v string) *GetSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSortScriptFileResponseBody) SetResult(v *GetSortScriptFileResponseBodyResult) *GetSortScriptFileResponseBody {
	s.Result = v
	return s
}

type GetSortScriptFileResponseBodyResult struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Version    *int64  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetSortScriptFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponseBodyResult) SetContent(v string) *GetSortScriptFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetCreateTime(v string) *GetSortScriptFileResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetModifyTime(v string) *GetSortScriptFileResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetVersion(v int64) *GetSortScriptFileResponseBodyResult {
	s.Version = &v
	return s
}

type GetSortScriptFileResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSortScriptFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponse) SetHeaders(v map[string]*string) *GetSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *GetSortScriptFileResponse) SetBody(v *GetSortScriptFileResponseBody) *GetSortScriptFileResponse {
	s.Body = v
	return s
}

type GetValidationErrorRequest struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
}

func (s GetValidationErrorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetValidationErrorRequest) GoString() string {
	return s.String()
}

func (s *GetValidationErrorRequest) SetErrorCode(v string) *GetValidationErrorRequest {
	s.ErrorCode = &v
	return s
}

type GetValidationErrorResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetValidationErrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetValidationErrorResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidationErrorResponseBody) SetRequestId(v string) *GetValidationErrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidationErrorResponseBody) SetResult(v []map[string]interface{}) *GetValidationErrorResponseBody {
	s.Result = v
	return s
}

type GetValidationErrorResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetValidationErrorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetValidationErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetValidationErrorResponse) GoString() string {
	return s.String()
}

func (s *GetValidationErrorResponse) SetHeaders(v map[string]*string) *GetValidationErrorResponse {
	s.Headers = v
	return s
}

func (s *GetValidationErrorResponse) SetBody(v *GetValidationErrorResponseBody) *GetValidationErrorResponse {
	s.Body = v
	return s
}

type GetValidationReportRequest struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetValidationReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetValidationReportRequest) GoString() string {
	return s.String()
}

func (s *GetValidationReportRequest) SetType(v string) *GetValidationReportRequest {
	s.Type = &v
	return s
}

type GetValidationReportResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetValidationReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetValidationReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidationReportResponseBody) SetRequestId(v string) *GetValidationReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidationReportResponseBody) SetResult(v []map[string]interface{}) *GetValidationReportResponseBody {
	s.Result = v
	return s
}

type GetValidationReportResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetValidationReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetValidationReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetValidationReportResponse) GoString() string {
	return s.String()
}

func (s *GetValidationReportResponse) SetHeaders(v map[string]*string) *GetValidationReportResponse {
	s.Headers = v
	return s
}

func (s *GetValidationReportResponse) SetBody(v *GetValidationReportResponseBody) *GetValidationReportResponse {
	s.Body = v
	return s
}

type ListABTestExperimentsResponseBody struct {
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListABTestExperimentsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponseBody) SetRequestId(v string) *ListABTestExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestExperimentsResponseBody) SetResult(v []*ListABTestExperimentsResponseBodyResult) *ListABTestExperimentsResponseBody {
	s.Result = v
	return s
}

type ListABTestExperimentsResponseBodyResult struct {
	Created *int32                 `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string                `json:"name,omitempty" xml:"name,omitempty"`
	Online  *bool                  `json:"online,omitempty" xml:"online,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Traffic *int32                 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	Updated *int32                 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListABTestExperimentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestExperimentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponseBodyResult) SetCreated(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetId(v string) *ListABTestExperimentsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetName(v string) *ListABTestExperimentsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetOnline(v bool) *ListABTestExperimentsResponseBodyResult {
	s.Online = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetParams(v map[string]interface{}) *ListABTestExperimentsResponseBodyResult {
	s.Params = v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetTraffic(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetUpdated(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListABTestExperimentsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListABTestExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponse) SetHeaders(v map[string]*string) *ListABTestExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestExperimentsResponse) SetBody(v *ListABTestExperimentsResponseBody) *ListABTestExperimentsResponse {
	s.Body = v
	return s
}

type ListABTestFixedFlowDividersResponseBody struct {
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestFixedFlowDividersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestFixedFlowDividersResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestFixedFlowDividersResponseBody) SetRequestId(v string) *ListABTestFixedFlowDividersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestFixedFlowDividersResponseBody) SetResult(v []*string) *ListABTestFixedFlowDividersResponseBody {
	s.Result = v
	return s
}

type ListABTestFixedFlowDividersResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestFixedFlowDividersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestFixedFlowDividersResponse) GoString() string {
	return s.String()
}

func (s *ListABTestFixedFlowDividersResponse) SetHeaders(v map[string]*string) *ListABTestFixedFlowDividersResponse {
	s.Headers = v
	return s
}

func (s *ListABTestFixedFlowDividersResponse) SetBody(v *ListABTestFixedFlowDividersResponseBody) *ListABTestFixedFlowDividersResponse {
	s.Body = v
	return s
}

type ListABTestGroupsResponseBody struct {
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListABTestGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponseBody) SetRequestId(v string) *ListABTestGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestGroupsResponseBody) SetResult(v []*ListABTestGroupsResponseBodyResult) *ListABTestGroupsResponseBody {
	s.Result = v
	return s
}

type ListABTestGroupsResponseBodyResult struct {
	Created *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListABTestGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponseBodyResult) SetCreated(v int32) *ListABTestGroupsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetId(v string) *ListABTestGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetName(v string) *ListABTestGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetStatus(v int32) *ListABTestGroupsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetUpdated(v int32) *ListABTestGroupsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListABTestGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListABTestGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponse) SetHeaders(v map[string]*string) *ListABTestGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestGroupsResponse) SetBody(v *ListABTestGroupsResponseBody) *ListABTestGroupsResponse {
	s.Body = v
	return s
}

type ListABTestMetricsResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListABTestMetricsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestMetricsResponseBody) SetRequestId(v string) *ListABTestMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestMetricsResponseBody) SetResult(v []*ListABTestMetricsResponseBodyResult) *ListABTestMetricsResponseBody {
	s.Result = v
	return s
}

type ListABTestMetricsResponseBodyResult struct {
	Ctr            *float32 `json:"ctr,omitempty" xml:"ctr,omitempty"`
	Date           *string  `json:"date,omitempty" xml:"date,omitempty"`
	ExperimentName *string  `json:"experimentName,omitempty" xml:"experimentName,omitempty"`
	Ipv            *int32   `json:"ipv,omitempty" xml:"ipv,omitempty"`
	IpvUv          *int32   `json:"ipvUv,omitempty" xml:"ipvUv,omitempty"`
	Pv             *int32   `json:"pv,omitempty" xml:"pv,omitempty"`
	Uv             *int32   `json:"uv,omitempty" xml:"uv,omitempty"`
	ZeroHitRate    *float32 `json:"zeroHitRate,omitempty" xml:"zeroHitRate,omitempty"`
}

func (s ListABTestMetricsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestMetricsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestMetricsResponseBodyResult) SetCtr(v float32) *ListABTestMetricsResponseBodyResult {
	s.Ctr = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetDate(v string) *ListABTestMetricsResponseBodyResult {
	s.Date = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetExperimentName(v string) *ListABTestMetricsResponseBodyResult {
	s.ExperimentName = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetIpv(v int32) *ListABTestMetricsResponseBodyResult {
	s.Ipv = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetIpvUv(v int32) *ListABTestMetricsResponseBodyResult {
	s.IpvUv = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetPv(v int32) *ListABTestMetricsResponseBodyResult {
	s.Pv = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetUv(v int32) *ListABTestMetricsResponseBodyResult {
	s.Uv = &v
	return s
}

func (s *ListABTestMetricsResponseBodyResult) SetZeroHitRate(v float32) *ListABTestMetricsResponseBodyResult {
	s.ZeroHitRate = &v
	return s
}

type ListABTestMetricsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListABTestMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestMetricsResponse) SetHeaders(v map[string]*string) *ListABTestMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestMetricsResponse) SetBody(v *ListABTestMetricsResponseBody) *ListABTestMetricsResponse {
	s.Body = v
	return s
}

type ListABTestScenesResponseBody struct {
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListABTestScenesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponseBody) SetRequestId(v string) *ListABTestScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestScenesResponseBody) SetResult(v []*ListABTestScenesResponseBodyResult) *ListABTestScenesResponseBody {
	s.Result = v
	return s
}

type ListABTestScenesResponseBodyResult struct {
	Created *int32    `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string   `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string   `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32    `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32    `json:"updated,omitempty" xml:"updated,omitempty"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListABTestScenesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestScenesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponseBodyResult) SetCreated(v int32) *ListABTestScenesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetId(v string) *ListABTestScenesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetName(v string) *ListABTestScenesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetStatus(v int32) *ListABTestScenesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetUpdated(v int32) *ListABTestScenesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetValues(v []*string) *ListABTestScenesResponseBodyResult {
	s.Values = v
	return s
}

type ListABTestScenesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListABTestScenesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestScenesResponse) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponse) SetHeaders(v map[string]*string) *ListABTestScenesResponse {
	s.Headers = v
	return s
}

func (s *ListABTestScenesResponse) SetBody(v *ListABTestScenesResponseBody) *ListABTestScenesResponse {
	s.Body = v
	return s
}

type ListAppGroupErrorsRequest struct {
	AppId      *string `json:"appId,omitempty" xml:"appId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime  *int32  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StopTime   *int32  `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
}

func (s ListAppGroupErrorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupErrorsRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupErrorsRequest) SetAppId(v string) *ListAppGroupErrorsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppGroupErrorsRequest) SetPageNumber(v int32) *ListAppGroupErrorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupErrorsRequest) SetPageSize(v int32) *ListAppGroupErrorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupErrorsRequest) SetStartTime(v int32) *ListAppGroupErrorsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAppGroupErrorsRequest) SetStopTime(v int32) *ListAppGroupErrorsRequest {
	s.StopTime = &v
	return s
}

type ListAppGroupErrorsResponseBody struct {
	RequestId  *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppGroupErrorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupErrorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupErrorsResponseBody) SetRequestId(v string) *ListAppGroupErrorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupErrorsResponseBody) SetResult(v []map[string]interface{}) *ListAppGroupErrorsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppGroupErrorsResponseBody) SetTotalCount(v int64) *ListAppGroupErrorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppGroupErrorsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppGroupErrorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppGroupErrorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupErrorsResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupErrorsResponse) SetHeaders(v map[string]*string) *ListAppGroupErrorsResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupErrorsResponse) SetBody(v *ListAppGroupErrorsResponseBody) *ListAppGroupErrorsResponse {
	s.Body = v
	return s
}

type ListAppGroupMetricsRequest struct {
	EndTime    *int32  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Indexes    *string `json:"indexes,omitempty" xml:"indexes,omitempty"`
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	StartTime  *int32  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListAppGroupMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupMetricsRequest) SetEndTime(v int32) *ListAppGroupMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAppGroupMetricsRequest) SetIndexes(v string) *ListAppGroupMetricsRequest {
	s.Indexes = &v
	return s
}

func (s *ListAppGroupMetricsRequest) SetMetricType(v string) *ListAppGroupMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *ListAppGroupMetricsRequest) SetStartTime(v int32) *ListAppGroupMetricsRequest {
	s.StartTime = &v
	return s
}

type ListAppGroupMetricsResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAppGroupMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupMetricsResponseBody) SetRequestId(v string) *ListAppGroupMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupMetricsResponseBody) SetResult(v []map[string]interface{}) *ListAppGroupMetricsResponseBody {
	s.Result = v
	return s
}

type ListAppGroupMetricsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppGroupMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppGroupMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupMetricsResponse) SetHeaders(v map[string]*string) *ListAppGroupMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupMetricsResponse) SetBody(v *ListAppGroupMetricsResponseBody) *ListAppGroupMetricsResponse {
	s.Body = v
	return s
}

type ListAppGroupsRequest struct {
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	PageNumber      *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	SortBy          *int32  `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAppGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupsRequest) SetInstanceId(v string) *ListAppGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsRequest) SetName(v string) *ListAppGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListAppGroupsRequest) SetPageNumber(v int32) *ListAppGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupsRequest) SetPageSize(v int32) *ListAppGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupsRequest) SetResourceGroupId(v string) *ListAppGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppGroupsRequest) SetSortBy(v int32) *ListAppGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListAppGroupsRequest) SetType(v string) *ListAppGroupsRequest {
	s.Type = &v
	return s
}

type ListAppGroupsResponseBody struct {
	RequestId  *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListAppGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBody) SetRequestId(v string) *ListAppGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupsResponseBody) SetResult(v []*ListAppGroupsResponseBodyResult) *ListAppGroupsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppGroupsResponseBody) SetTotalCount(v int32) *ListAppGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppGroupsResponseBodyResult struct {
	ChargeType                        *string                               `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay                       *int32                                `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	CommodityCode                     *string                               `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Created                           *int32                                `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion                    *string                               `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Description                       *string                               `json:"description,omitempty" xml:"description,omitempty"`
	Domain                            *string                               `json:"domain,omitempty" xml:"domain,omitempty"`
	ExpireOn                          *string                               `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	FirstRankAlgoDeploymentId         *int32                                `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	HasPendingQuotaReviewTask         *int32                                `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	Id                                *string                               `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId                        *string                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockMode                          *string                               `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration                *int32                                `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	Name                              *string                               `json:"name,omitempty" xml:"name,omitempty"`
	PendingSecondRankAlgoDeploymentId *int32                                `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	ProcessingOrderId                 *string                               `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	Produced                          *int32                                `json:"produced,omitempty" xml:"produced,omitempty"`
	ProjectId                         *string                               `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Quota                             *ListAppGroupsResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	ResourceGroupId                   *string                               `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	SecondRankAlgoDeploymentId        *int32                                `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	Status                            *string                               `json:"status,omitempty" xml:"status,omitempty"`
	SwitchedTime                      *int32                                `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Type                              *string                               `json:"type,omitempty" xml:"type,omitempty"`
	Updated                           *int32                                `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListAppGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResult) SetChargeType(v string) *ListAppGroupsResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetChargingWay(v int32) *ListAppGroupsResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCommodityCode(v string) *ListAppGroupsResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCreated(v int32) *ListAppGroupsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCurrentVersion(v string) *ListAppGroupsResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetDescription(v string) *ListAppGroupsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetDomain(v string) *ListAppGroupsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetExpireOn(v string) *ListAppGroupsResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ListAppGroupsResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ListAppGroupsResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetId(v string) *ListAppGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetInstanceId(v string) *ListAppGroupsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetLockMode(v string) *ListAppGroupsResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetLockedByExpiration(v int32) *ListAppGroupsResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetName(v string) *ListAppGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ListAppGroupsResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProcessingOrderId(v string) *ListAppGroupsResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProduced(v int32) *ListAppGroupsResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProjectId(v string) *ListAppGroupsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetQuota(v *ListAppGroupsResponseBodyResultQuota) *ListAppGroupsResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetResourceGroupId(v string) *ListAppGroupsResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ListAppGroupsResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetStatus(v string) *ListAppGroupsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetSwitchedTime(v int32) *ListAppGroupsResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetType(v string) *ListAppGroupsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetUpdated(v int32) *ListAppGroupsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListAppGroupsResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListAppGroupsResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResultQuota) SetComputeResource(v int32) *ListAppGroupsResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) SetDocSize(v int32) *ListAppGroupsResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) SetSpec(v string) *ListAppGroupsResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ListAppGroupsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponse) SetHeaders(v map[string]*string) *ListAppGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupsResponse) SetBody(v *ListAppGroupsResponseBody) *ListAppGroupsResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	Group *bool  `json:"group,omitempty" xml:"group,omitempty"`
	Page  *int32 `json:"page,omitempty" xml:"page,omitempty"`
	Size  *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetGroup(v bool) *ListAppsRequest {
	s.Group = &v
	return s
}

func (s *ListAppsRequest) SetPage(v int32) *ListAppsRequest {
	s.Page = &v
	return s
}

func (s *ListAppsRequest) SetSize(v int32) *ListAppsRequest {
	s.Size = &v
	return s
}

type ListAppsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

type ListDataCollectionsRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDataCollectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsRequest) SetPageNumber(v int32) *ListDataCollectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCollectionsRequest) SetPageSize(v int32) *ListDataCollectionsRequest {
	s.PageSize = &v
	return s
}

type ListDataCollectionsResponseBody struct {
	RequestId  *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListDataCollectionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCollectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponseBody) SetRequestId(v string) *ListDataCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCollectionsResponseBody) SetResult(v []*ListDataCollectionsResponseBodyResult) *ListDataCollectionsResponseBody {
	s.Result = v
	return s
}

func (s *ListDataCollectionsResponseBody) SetTotalCount(v int32) *ListDataCollectionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataCollectionsResponseBodyResult struct {
	Created            *int32  `json:"created,omitempty" xml:"created,omitempty"`
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	Id                 *string `json:"id,omitempty" xml:"id,omitempty"`
	IndustryName       *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Status             *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SundialId          *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated            *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListDataCollectionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponseBodyResult) SetCreated(v int32) *ListDataCollectionsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetDataCollectionType(v string) *ListDataCollectionsResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetId(v string) *ListDataCollectionsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetIndustryName(v string) *ListDataCollectionsResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetName(v string) *ListDataCollectionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetStatus(v int32) *ListDataCollectionsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetSundialId(v string) *ListDataCollectionsResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetType(v string) *ListDataCollectionsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetUpdated(v int32) *ListDataCollectionsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListDataCollectionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataCollectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponse) SetHeaders(v map[string]*string) *ListDataCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataCollectionsResponse) SetBody(v *ListDataCollectionsResponseBody) *ListDataCollectionsResponse {
	s.Body = v
	return s
}

type ListDataSourceTableFieldsRequest struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ListDataSourceTableFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTableFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsRequest) SetParams(v string) *ListDataSourceTableFieldsRequest {
	s.Params = &v
	return s
}

type ListDataSourceTableFieldsResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListDataSourceTableFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTableFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsResponseBody) SetRequestId(v string) *ListDataSourceTableFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTableFieldsResponseBody) SetResult(v map[string]interface{}) *ListDataSourceTableFieldsResponseBody {
	s.Result = v
	return s
}

type ListDataSourceTableFieldsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourceTableFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceTableFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTableFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsResponse) SetHeaders(v map[string]*string) *ListDataSourceTableFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTableFieldsResponse) SetBody(v *ListDataSourceTableFieldsResponseBody) *ListDataSourceTableFieldsResponse {
	s.Body = v
	return s
}

type ListDataSourceTablesRequest struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ListDataSourceTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesRequest) SetParams(v string) *ListDataSourceTablesRequest {
	s.Params = &v
	return s
}

type ListDataSourceTablesResponseBody struct {
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesResponseBody) SetRequestId(v string) *ListDataSourceTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTablesResponseBody) SetResult(v []*string) *ListDataSourceTablesResponseBody {
	s.Result = v
	return s
}

type ListDataSourceTablesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesResponse) SetHeaders(v map[string]*string) *ListDataSourceTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTablesResponse) SetBody(v *ListDataSourceTablesResponseBody) *ListDataSourceTablesResponse {
	s.Body = v
	return s
}

type ListDeployedAlgorithmModelsRequest struct {
	AlgorithmType *string `json:"algorithmType,omitempty" xml:"algorithmType,omitempty"`
	InServiceOnly *bool   `json:"inServiceOnly,omitempty" xml:"inServiceOnly,omitempty"`
}

func (s ListDeployedAlgorithmModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedAlgorithmModelsRequest) GoString() string {
	return s.String()
}

func (s *ListDeployedAlgorithmModelsRequest) SetAlgorithmType(v string) *ListDeployedAlgorithmModelsRequest {
	s.AlgorithmType = &v
	return s
}

func (s *ListDeployedAlgorithmModelsRequest) SetInServiceOnly(v bool) *ListDeployedAlgorithmModelsRequest {
	s.InServiceOnly = &v
	return s
}

type ListDeployedAlgorithmModelsResponseBody struct {
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListDeployedAlgorithmModelsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDeployedAlgorithmModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedAlgorithmModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployedAlgorithmModelsResponseBody) SetRequestId(v string) *ListDeployedAlgorithmModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBody) SetResult(v []*ListDeployedAlgorithmModelsResponseBodyResult) *ListDeployedAlgorithmModelsResponseBody {
	s.Result = v
	return s
}

type ListDeployedAlgorithmModelsResponseBodyResult struct {
	AppGroupName *string                                                `json:"appGroupName,omitempty" xml:"appGroupName,omitempty"`
	Apps         []*string                                              `json:"apps,omitempty" xml:"apps,omitempty" type:"Repeated"`
	Desc         *string                                                `json:"desc,omitempty" xml:"desc,omitempty"`
	GmtCreate    *string                                                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string                                                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id           *string                                                `json:"id,omitempty" xml:"id,omitempty"`
	Models       []*ListDeployedAlgorithmModelsResponseBodyResultModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	Scene        *string                                                `json:"scene,omitempty" xml:"scene,omitempty"`
	Status       *string                                                `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDeployedAlgorithmModelsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedAlgorithmModelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetAppGroupName(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.AppGroupName = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetApps(v []*string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.Apps = v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetDesc(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetGmtCreate(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetGmtModified(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetId(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetModels(v []*ListDeployedAlgorithmModelsResponseBodyResultModels) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.Models = v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetScene(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.Scene = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResult) SetStatus(v string) *ListDeployedAlgorithmModelsResponseBodyResult {
	s.Status = &v
	return s
}

type ListDeployedAlgorithmModelsResponseBodyResultModels struct {
	AlgorithmType *string `json:"algorithmType,omitempty" xml:"algorithmType,omitempty"`
	ModelId       *int32  `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelName     *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	Progress      *int32  `json:"progress,omitempty" xml:"progress,omitempty"`
	ProjectId     *int32  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDeployedAlgorithmModelsResponseBodyResultModels) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedAlgorithmModelsResponseBodyResultModels) GoString() string {
	return s.String()
}

func (s *ListDeployedAlgorithmModelsResponseBodyResultModels) SetAlgorithmType(v string) *ListDeployedAlgorithmModelsResponseBodyResultModels {
	s.AlgorithmType = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResultModels) SetModelId(v int32) *ListDeployedAlgorithmModelsResponseBodyResultModels {
	s.ModelId = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResultModels) SetModelName(v string) *ListDeployedAlgorithmModelsResponseBodyResultModels {
	s.ModelName = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResultModels) SetProgress(v int32) *ListDeployedAlgorithmModelsResponseBodyResultModels {
	s.Progress = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResultModels) SetProjectId(v int32) *ListDeployedAlgorithmModelsResponseBodyResultModels {
	s.ProjectId = &v
	return s
}

func (s *ListDeployedAlgorithmModelsResponseBodyResultModels) SetStatus(v string) *ListDeployedAlgorithmModelsResponseBodyResultModels {
	s.Status = &v
	return s
}

type ListDeployedAlgorithmModelsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeployedAlgorithmModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployedAlgorithmModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployedAlgorithmModelsResponse) GoString() string {
	return s.String()
}

func (s *ListDeployedAlgorithmModelsResponse) SetHeaders(v map[string]*string) *ListDeployedAlgorithmModelsResponse {
	s.Headers = v
	return s
}

func (s *ListDeployedAlgorithmModelsResponse) SetBody(v *ListDeployedAlgorithmModelsResponseBody) *ListDeployedAlgorithmModelsResponse {
	s.Body = v
	return s
}

type ListFirstRanksResponseBody struct {
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListFirstRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListFirstRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBody) SetRequestId(v string) *ListFirstRanksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFirstRanksResponseBody) SetResult(v []*ListFirstRanksResponseBodyResult) *ListFirstRanksResponseBody {
	s.Result = v
	return s
}

type ListFirstRanksResponseBodyResult struct {
	Active      *bool                                   `json:"active,omitempty" xml:"active,omitempty"`
	Created     *int32                                  `json:"created,omitempty" xml:"created,omitempty"`
	Description *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	Meta        []*ListFirstRanksResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name        *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Updated     *int32                                  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListFirstRanksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBodyResult) SetActive(v bool) *ListFirstRanksResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetCreated(v int32) *ListFirstRanksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetDescription(v string) *ListFirstRanksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetMeta(v []*ListFirstRanksResponseBodyResultMeta) *ListFirstRanksResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetName(v string) *ListFirstRanksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetUpdated(v int32) *ListFirstRanksResponseBodyResult {
	s.Updated = &v
	return s
}

type ListFirstRanksResponseBodyResultMeta struct {
	Arg       *string `json:"arg,omitempty" xml:"arg,omitempty"`
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Weight    *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListFirstRanksResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBodyResultMeta) SetArg(v string) *ListFirstRanksResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) SetAttribute(v string) *ListFirstRanksResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) SetWeight(v int32) *ListFirstRanksResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type ListFirstRanksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFirstRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFirstRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponse) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponse) SetHeaders(v map[string]*string) *ListFirstRanksResponse {
	s.Headers = v
	return s
}

func (s *ListFirstRanksResponse) SetBody(v *ListFirstRanksResponseBody) *ListFirstRanksResponse {
	s.Body = v
	return s
}

type ListFunctionInstancesRequest struct {
	// 功能类型
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// 模型类型
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// 返回信息的丰富度
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// 页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 实例来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListFunctionInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesRequest) SetFunctionType(v string) *ListFunctionInstancesRequest {
	s.FunctionType = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetModelType(v string) *ListFunctionInstancesRequest {
	s.ModelType = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetOutput(v string) *ListFunctionInstancesRequest {
	s.Output = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetPageNumber(v int32) *ListFunctionInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetPageSize(v int32) *ListFunctionInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetSource(v string) *ListFunctionInstancesRequest {
	s.Source = &v
	return s
}

type ListFunctionInstancesResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode   *int64                                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency    *int64                                     `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListFunctionInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Status     *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBody) SetCode(v string) *ListFunctionInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetHttpCode(v int64) *ListFunctionInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetLatency(v int64) *ListFunctionInstancesResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetMessage(v string) *ListFunctionInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetRequestId(v string) *ListFunctionInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetResult(v []*ListFunctionInstancesResponseBodyResult) *ListFunctionInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetStatus(v string) *ListFunctionInstancesResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetTotalCount(v int64) *ListFunctionInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFunctionInstancesResponseBodyResult struct {
	Belongs          *ListFunctionInstancesResponseBodyResultBelongs            `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	CreateParameters []*ListFunctionInstancesResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	CreateTime       *int64                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Cron             *string                                                    `json:"Cron,omitempty" xml:"Cron,omitempty"`
	Description      *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	ExtendInfo       *string                                                    `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FunctionName     *string                                                    `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionType     *string                                                    `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	InstanceName     *string                                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ModelType        *string                                                    `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	Source           *string                                                    `json:"Source,omitempty" xml:"Source,omitempty"`
	Status           *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionId        *int64                                                     `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResult) SetBelongs(v *ListFunctionInstancesResponseBodyResultBelongs) *ListFunctionInstancesResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCreateParameters(v []*ListFunctionInstancesResponseBodyResultCreateParameters) *ListFunctionInstancesResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCreateTime(v int64) *ListFunctionInstancesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCron(v string) *ListFunctionInstancesResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetDescription(v string) *ListFunctionInstancesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetExtendInfo(v string) *ListFunctionInstancesResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetFunctionName(v string) *ListFunctionInstancesResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetFunctionType(v string) *ListFunctionInstancesResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetInstanceName(v string) *ListFunctionInstancesResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetModelType(v string) *ListFunctionInstancesResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetSource(v string) *ListFunctionInstancesResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetStatus(v string) *ListFunctionInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetVersionId(v int64) *ListFunctionInstancesResponseBodyResult {
	s.VersionId = &v
	return s
}

type ListFunctionInstancesResponseBodyResultBelongs struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultBelongs) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetCategory(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetDomain(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetLanguage(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Language = &v
	return s
}

type ListFunctionInstancesResponseBodyResultCreateParameters struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) SetName(v string) *ListFunctionInstancesResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) SetValue(v string) *ListFunctionInstancesResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

type ListFunctionInstancesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponse) SetHeaders(v map[string]*string) *ListFunctionInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionInstancesResponse) SetBody(v *ListFunctionInstancesResponseBody) *ListFunctionInstancesResponse {
	s.Body = v
	return s
}

type ListFunctionTasksRequest struct {
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListFunctionTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksRequest) SetEndTime(v int64) *ListFunctionTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListFunctionTasksRequest) SetPageNumber(v int32) *ListFunctionTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionTasksRequest) SetPageSize(v int32) *ListFunctionTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionTasksRequest) SetStartTime(v int64) *ListFunctionTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListFunctionTasksRequest) SetStatus(v string) *ListFunctionTasksRequest {
	s.Status = &v
	return s
}

type ListFunctionTasksResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode   *int64                                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency    *int64                                 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListFunctionTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Status     *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponseBody) SetCode(v string) *ListFunctionTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetHttpCode(v int64) *ListFunctionTasksResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetLatency(v int64) *ListFunctionTasksResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetMessage(v string) *ListFunctionTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetRequestId(v string) *ListFunctionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetResult(v []*ListFunctionTasksResponseBodyResult) *ListFunctionTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionTasksResponseBody) SetStatus(v string) *ListFunctionTasksResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetTotalCount(v int64) *ListFunctionTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListFunctionTasksResponseBodyResult struct {
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtendInfo   *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Progress     *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RunId        *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponseBodyResult) SetEndTime(v int64) *ListFunctionTasksResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetExtendInfo(v string) *ListFunctionTasksResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetFunctionName(v string) *ListFunctionTasksResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetProgress(v int64) *ListFunctionTasksResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetRunId(v string) *ListFunctionTasksResponseBodyResult {
	s.RunId = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetStartTime(v int64) *ListFunctionTasksResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetStatus(v string) *ListFunctionTasksResponseBodyResult {
	s.Status = &v
	return s
}

type ListFunctionTasksResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponse) SetHeaders(v map[string]*string) *ListFunctionTasksResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionTasksResponse) SetBody(v *ListFunctionTasksResponseBody) *ListFunctionTasksResponse {
	s.Body = v
	return s
}

type ListInterventionDictionariesRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Types      *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListInterventionDictionariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesRequest) SetPageNumber(v int32) *ListInterventionDictionariesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInterventionDictionariesRequest) SetPageSize(v int32) *ListInterventionDictionariesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterventionDictionariesRequest) SetTypes(v string) *ListInterventionDictionariesRequest {
	s.Types = &v
	return s
}

type ListInterventionDictionariesResponseBody struct {
	RequestId  *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListInterventionDictionariesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInterventionDictionariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponseBody) SetRequestId(v string) *ListInterventionDictionariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionariesResponseBody) SetResult(v []*ListInterventionDictionariesResponseBodyResult) *ListInterventionDictionariesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionariesResponseBody) SetTotalCount(v int32) *ListInterventionDictionariesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInterventionDictionariesResponseBodyResult struct {
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	Created  *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Id       *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated  *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListInterventionDictionariesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponseBodyResult) SetAnalyzer(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetCreated(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetId(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetName(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetType(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetUpdated(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Updated = &v
	return s
}

type ListInterventionDictionariesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInterventionDictionariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionariesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionariesResponse) SetBody(v *ListInterventionDictionariesResponseBody) *ListInterventionDictionariesResponse {
	s.Body = v
	return s
}

type ListInterventionDictionaryEntriesRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Word       *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListInterventionDictionaryEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesRequest) SetPageNumber(v int32) *ListInterventionDictionaryEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) SetPageSize(v int32) *ListInterventionDictionaryEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) SetWord(v string) *ListInterventionDictionaryEntriesRequest {
	s.Word = &v
	return s
}

type ListInterventionDictionaryEntriesResponseBody struct {
	RequestId  *string                                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListInterventionDictionaryEntriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetRequestId(v string) *ListInterventionDictionaryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetResult(v []*ListInterventionDictionaryEntriesResponseBodyResult) *ListInterventionDictionaryEntriesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetTotalCount(v int32) *ListInterventionDictionaryEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInterventionDictionaryEntriesResponseBodyResult struct {
	Cmd       *string                                                      `json:"cmd,omitempty" xml:"cmd,omitempty"`
	Created   *int64                                                       `json:"created,omitempty" xml:"created,omitempty"`
	Relevance map[string]interface{}                                       `json:"relevance,omitempty" xml:"relevance,omitempty"`
	Status    *string                                                      `json:"status,omitempty" xml:"status,omitempty"`
	Tokens    []*ListInterventionDictionaryEntriesResponseBodyResultTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
	Updated   *int64                                                       `json:"updated,omitempty" xml:"updated,omitempty"`
	Word      *string                                                      `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetCmd(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Cmd = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetCreated(v int64) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetRelevance(v map[string]interface{}) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Relevance = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetStatus(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetTokens(v []*ListInterventionDictionaryEntriesResponseBodyResultTokens) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Tokens = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetUpdated(v int64) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetWord(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Word = &v
	return s
}

type ListInterventionDictionaryEntriesResponseBodyResultTokens struct {
	Order    *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Tag      *string `json:"tag,omitempty" xml:"tag,omitempty"`
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	Token    *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBodyResultTokens) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBodyResultTokens) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetOrder(v int32) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Order = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetTag(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Tag = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetTagLabel(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.TagLabel = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetToken(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Token = &v
	return s
}

type ListInterventionDictionaryEntriesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionaryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponse) SetBody(v *ListInterventionDictionaryEntriesResponseBody) *ListInterventionDictionaryEntriesResponse {
	s.Body = v
	return s
}

type ListInterventionDictionaryNerResultsRequest struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListInterventionDictionaryNerResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsRequest) SetQuery(v string) *ListInterventionDictionaryNerResultsRequest {
	s.Query = &v
	return s
}

type ListInterventionDictionaryNerResultsResponseBody struct {
	RequestId *string                                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListInterventionDictionaryNerResultsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInterventionDictionaryNerResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponseBody) SetRequestId(v string) *ListInterventionDictionaryNerResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBody) SetResult(v []*ListInterventionDictionaryNerResultsResponseBodyResult) *ListInterventionDictionaryNerResultsResponseBody {
	s.Result = v
	return s
}

type ListInterventionDictionaryNerResultsResponseBodyResult struct {
	Order    *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Tag      *string `json:"tag,omitempty" xml:"tag,omitempty"`
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	Token    *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListInterventionDictionaryNerResultsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetOrder(v int32) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetTag(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Tag = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetTagLabel(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.TagLabel = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetToken(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Token = &v
	return s
}

type ListInterventionDictionaryNerResultsResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInterventionDictionaryNerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionaryNerResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryNerResultsResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponse) SetBody(v *ListInterventionDictionaryNerResultsResponseBody) *ListInterventionDictionaryNerResultsResponse {
	s.Body = v
	return s
}

type ListInterventionDictionaryRelatedEntitiesResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInterventionDictionaryRelatedEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryRelatedEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) SetRequestId(v string) *ListInterventionDictionaryRelatedEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) SetResult(v []map[string]interface{}) *ListInterventionDictionaryRelatedEntitiesResponseBody {
	s.Result = v
	return s
}

type ListInterventionDictionaryRelatedEntitiesResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInterventionDictionaryRelatedEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionaryRelatedEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryRelatedEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetBody(v *ListInterventionDictionaryRelatedEntitiesResponseBody) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.Body = v
	return s
}

type ListModelsRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) SetPageNumber(v int32) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int32) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) SetType(v string) *ListModelsRequest {
	s.Type = &v
	return s
}

type ListModelsResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetResult(v []map[string]interface{}) *ListModelsResponseBody {
	s.Result = v
	return s
}

type ListModelsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponse) SetHeaders(v map[string]*string) *ListModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelsResponse) SetBody(v *ListModelsResponseBody) *ListModelsResponse {
	s.Body = v
	return s
}

type ListProceedingsResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProceedingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProceedingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProceedingsResponseBody) SetRequestId(v string) *ListProceedingsResponseBody {
	s.RequestId = &v
	return s
}

type ListProceedingsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProceedingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProceedingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProceedingsResponse) GoString() string {
	return s.String()
}

func (s *ListProceedingsResponse) SetHeaders(v map[string]*string) *ListProceedingsResponse {
	s.Headers = v
	return s
}

func (s *ListProceedingsResponse) SetBody(v *ListProceedingsResponseBody) *ListProceedingsResponse {
	s.Body = v
	return s
}

type ListQueryProcessorAnalyzerResultsRequest struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsRequest) SetText(v string) *ListQueryProcessorAnalyzerResultsRequest {
	s.Text = &v
	return s
}

type ListQueryProcessorAnalyzerResultsResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) SetRequestId(v string) *ListQueryProcessorAnalyzerResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) SetResult(v map[string]interface{}) *ListQueryProcessorAnalyzerResultsResponseBody {
	s.Result = v
	return s
}

type ListQueryProcessorAnalyzerResultsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueryProcessorAnalyzerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryProcessorAnalyzerResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetHeaders(v map[string]*string) *ListQueryProcessorAnalyzerResultsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetBody(v *ListQueryProcessorAnalyzerResultsResponseBody) *ListQueryProcessorAnalyzerResultsResponse {
	s.Body = v
	return s
}

type ListQueryProcessorNersRequest struct {
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s ListQueryProcessorNersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersRequest) SetDomain(v string) *ListQueryProcessorNersRequest {
	s.Domain = &v
	return s
}

type ListQueryProcessorNersResponseBody struct {
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListQueryProcessorNersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListQueryProcessorNersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponseBody) SetRequestId(v string) *ListQueryProcessorNersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorNersResponseBody) SetResult(v []*ListQueryProcessorNersResponseBodyResult) *ListQueryProcessorNersResponseBody {
	s.Result = v
	return s
}

type ListQueryProcessorNersResponseBodyResult struct {
	Label    *string `json:"label,omitempty" xml:"label,omitempty"`
	Order    *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	Tag      *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListQueryProcessorNersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponseBodyResult) SetLabel(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Label = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetOrder(v int32) *ListQueryProcessorNersResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetPriority(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetTag(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Tag = &v
	return s
}

type ListQueryProcessorNersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueryProcessorNersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryProcessorNersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponse) SetHeaders(v map[string]*string) *ListQueryProcessorNersResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorNersResponse) SetBody(v *ListQueryProcessorNersResponseBody) *ListQueryProcessorNersResponse {
	s.Body = v
	return s
}

type ListQueryProcessorsRequest struct {
	IsActive *int32 `json:"isActive,omitempty" xml:"isActive,omitempty"`
}

func (s ListQueryProcessorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsRequest) SetIsActive(v int32) *ListQueryProcessorsRequest {
	s.IsActive = &v
	return s
}

type ListQueryProcessorsResponseBody struct {
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListQueryProcessorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListQueryProcessorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponseBody) SetRequestId(v string) *ListQueryProcessorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorsResponseBody) SetResult(v []*ListQueryProcessorsResponseBodyResult) *ListQueryProcessorsResponseBody {
	s.Result = v
	return s
}

type ListQueryProcessorsResponseBodyResult struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Created    *int32                   `json:"created,omitempty" xml:"created,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	Updated    *int32                   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListQueryProcessorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponseBodyResult) SetActive(v bool) *ListQueryProcessorsResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetCreated(v int32) *ListQueryProcessorsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetDomain(v string) *ListQueryProcessorsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetIndexes(v []*string) *ListQueryProcessorsResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetName(v string) *ListQueryProcessorsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetProcessors(v []map[string]interface{}) *ListQueryProcessorsResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetUpdated(v int32) *ListQueryProcessorsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListQueryProcessorsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueryProcessorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryProcessorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponse) SetHeaders(v map[string]*string) *ListQueryProcessorsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorsResponse) SetBody(v *ListQueryProcessorsResponseBody) *ListQueryProcessorsResponse {
	s.Body = v
	return s
}

type ListQuotaReviewTasksRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListQuotaReviewTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksRequest) SetPageNumber(v int32) *ListQuotaReviewTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotaReviewTasksRequest) SetPageSize(v int32) *ListQuotaReviewTasksRequest {
	s.PageSize = &v
	return s
}

type ListQuotaReviewTasksResponseBody struct {
	RequestId  *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListQuotaReviewTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListQuotaReviewTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponseBody) SetRequestId(v string) *ListQuotaReviewTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) SetResult(v []*ListQuotaReviewTasksResponseBodyResult) *ListQuotaReviewTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) SetTotalCount(v int32) *ListQuotaReviewTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaReviewTasksResponseBodyResult struct {
	AppGroupId         *int32  `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	AppGroupName       *string `json:"appGroupName,omitempty" xml:"appGroupName,omitempty"`
	AppGroupType       *string `json:"appGroupType,omitempty" xml:"appGroupType,omitempty"`
	Approved           *bool   `json:"approved,omitempty" xml:"approved,omitempty"`
	Available          *bool   `json:"available,omitempty" xml:"available,omitempty"`
	GmtCreate          *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                 *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Memo               *string `json:"memo,omitempty" xml:"memo,omitempty"`
	NewComputeResource *int32  `json:"newComputeResource,omitempty" xml:"newComputeResource,omitempty"`
	NewSocSize         *int32  `json:"newSocSize,omitempty" xml:"newSocSize,omitempty"`
	NewSpec            *string `json:"newSpec,omitempty" xml:"newSpec,omitempty"`
	OldComputeResource *int32  `json:"oldComputeResource,omitempty" xml:"oldComputeResource,omitempty"`
	OldDocSize         *int32  `json:"oldDocSize,omitempty" xml:"oldDocSize,omitempty"`
	OldSpec            *string `json:"oldSpec,omitempty" xml:"oldSpec,omitempty"`
	Pending            *bool   `json:"pending,omitempty" xml:"pending,omitempty"`
}

func (s ListQuotaReviewTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupId(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupName(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupName = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupType(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupType = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetApproved(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Approved = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAvailable(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetGmtCreate(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetGmtModified(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetId(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetMemo(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewComputeResource(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.NewComputeResource = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewSocSize(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.NewSocSize = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewSpec(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.NewSpec = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldComputeResource(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.OldComputeResource = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldDocSize(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.OldDocSize = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldSpec(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.OldSpec = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetPending(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Pending = &v
	return s
}

type ListQuotaReviewTasksResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQuotaReviewTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQuotaReviewTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponse) SetHeaders(v map[string]*string) *ListQuotaReviewTasksResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaReviewTasksResponse) SetBody(v *ListQuotaReviewTasksResponseBody) *ListQuotaReviewTasksResponse {
	s.Body = v
	return s
}

type ListRamRolesResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRamRolesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListRamRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRamRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRamRolesResponseBody) SetRequestId(v string) *ListRamRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRamRolesResponseBody) SetResult(v []*ListRamRolesResponseBodyResult) *ListRamRolesResponseBody {
	s.Result = v
	return s
}

type ListRamRolesResponseBodyResult struct {
	Assumed    *bool   `json:"assumed,omitempty" xml:"assumed,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Service    *string `json:"service,omitempty" xml:"service,omitempty"`
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s ListRamRolesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRamRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRamRolesResponseBodyResult) SetAssumed(v bool) *ListRamRolesResponseBodyResult {
	s.Assumed = &v
	return s
}

func (s *ListRamRolesResponseBodyResult) SetName(v string) *ListRamRolesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRamRolesResponseBodyResult) SetService(v string) *ListRamRolesResponseBodyResult {
	s.Service = &v
	return s
}

func (s *ListRamRolesResponseBodyResult) SetTemplateId(v string) *ListRamRolesResponseBodyResult {
	s.TemplateId = &v
	return s
}

type ListRamRolesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRamRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRamRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRamRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRamRolesResponse) SetHeaders(v map[string]*string) *ListRamRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRamRolesResponse) SetBody(v *ListRamRolesResponseBody) *ListRamRolesResponse {
	s.Body = v
	return s
}

type ListScheduledTasksRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListScheduledTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksRequest) SetPageNumber(v int32) *ListScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPageSize(v int32) *ListScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksRequest) SetType(v string) *ListScheduledTasksRequest {
	s.Type = &v
	return s
}

type ListScheduledTasksResponseBody struct {
	RequestId  *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListScheduledTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBody) SetRequestId(v string) *ListScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetResult(v []map[string]interface{}) *ListScheduledTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListScheduledTasksResponseBody) SetTotalCount(v int64) *ListScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListScheduledTasksResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScheduledTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledTasksResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponse) SetHeaders(v map[string]*string) *ListScheduledTasksResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledTasksResponse) SetBody(v *ListScheduledTasksResponseBody) *ListScheduledTasksResponse {
	s.Body = v
	return s
}

type ListSearchStrategiesResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSearchStrategiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSearchStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchStrategiesResponseBody) SetRequestId(v string) *ListSearchStrategiesResponseBody {
	s.RequestId = &v
	return s
}

type ListSearchStrategiesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSearchStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSearchStrategiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSearchStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListSearchStrategiesResponse) SetHeaders(v map[string]*string) *ListSearchStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListSearchStrategiesResponse) SetBody(v *ListSearchStrategiesResponseBody) *ListSearchStrategiesResponse {
	s.Body = v
	return s
}

type ListSecondRanksResponseBody struct {
	RequestId  *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListSecondRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSecondRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecondRanksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponseBody) SetRequestId(v string) *ListSecondRanksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecondRanksResponseBody) SetResult(v []*ListSecondRanksResponseBodyResult) *ListSecondRanksResponseBody {
	s.Result = v
	return s
}

func (s *ListSecondRanksResponseBody) SetTotalCount(v int32) *ListSecondRanksResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecondRanksResponseBodyResult struct {
	Active      *bool   `json:"active,omitempty" xml:"active,omitempty"`
	Created     *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDefault   *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	IsSys       *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	Meta        *string `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated     *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListSecondRanksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSecondRanksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponseBodyResult) SetActive(v bool) *ListSecondRanksResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetCreated(v int32) *ListSecondRanksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetDescription(v string) *ListSecondRanksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetId(v string) *ListSecondRanksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetIsDefault(v string) *ListSecondRanksResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetIsSys(v string) *ListSecondRanksResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetMeta(v string) *ListSecondRanksResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetName(v string) *ListSecondRanksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetUpdated(v int32) *ListSecondRanksResponseBodyResult {
	s.Updated = &v
	return s
}

type ListSecondRanksResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecondRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecondRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecondRanksResponse) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponse) SetHeaders(v map[string]*string) *ListSecondRanksResponse {
	s.Headers = v
	return s
}

func (s *ListSecondRanksResponse) SetBody(v *ListSecondRanksResponseBody) *ListSecondRanksResponse {
	s.Body = v
	return s
}

type ListSlowQueryCategoriesResponseBody struct {
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ListSlowQueryCategoriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSlowQueryCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponseBody) SetRequestId(v string) *ListSlowQueryCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBody) SetResult(v *ListSlowQueryCategoriesResponseBodyResult) *ListSlowQueryCategoriesResponseBody {
	s.Result = v
	return s
}

type ListSlowQueryCategoriesResponseBodyResult struct {
	AnalyzeStatus *string `json:"analyzeStatus,omitempty" xml:"analyzeStatus,omitempty"`
	End           *int32  `json:"end,omitempty" xml:"end,omitempty"`
	Start         *int32  `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListSlowQueryCategoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryCategoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetAnalyzeStatus(v string) *ListSlowQueryCategoriesResponseBodyResult {
	s.AnalyzeStatus = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetEnd(v int32) *ListSlowQueryCategoriesResponseBodyResult {
	s.End = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetStart(v int32) *ListSlowQueryCategoriesResponseBodyResult {
	s.Start = &v
	return s
}

type ListSlowQueryCategoriesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSlowQueryCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlowQueryCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponse) SetHeaders(v map[string]*string) *ListSlowQueryCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListSlowQueryCategoriesResponse) SetBody(v *ListSlowQueryCategoriesResponseBody) *ListSlowQueryCategoriesResponse {
	s.Body = v
	return s
}

type ListSlowQueryQueriesResponseBody struct {
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ListSlowQueryQueriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSlowQueryQueriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponseBody) SetRequestId(v string) *ListSlowQueryQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBody) SetResult(v *ListSlowQueryQueriesResponseBodyResult) *ListSlowQueryQueriesResponseBody {
	s.Result = v
	return s
}

type ListSlowQueryQueriesResponseBodyResult struct {
	AppQuery *string `json:"appQuery,omitempty" xml:"appQuery,omitempty"`
	End      *int32  `json:"end,omitempty" xml:"end,omitempty"`
	Index    *int32  `json:"index,omitempty" xml:"index,omitempty"`
	Start    *int32  `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListSlowQueryQueriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryQueriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetAppQuery(v string) *ListSlowQueryQueriesResponseBodyResult {
	s.AppQuery = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetEnd(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.End = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetIndex(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetStart(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.Start = &v
	return s
}

type ListSlowQueryQueriesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSlowQueryQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlowQueryQueriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponse) SetHeaders(v map[string]*string) *ListSlowQueryQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListSlowQueryQueriesResponse) SetBody(v *ListSlowQueryQueriesResponseBody) *ListSlowQueryQueriesResponse {
	s.Body = v
	return s
}

type ListSortExpressionsResponseBody struct {
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSortExpressionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSortExpressionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSortExpressionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponseBody) SetRequestId(v string) *ListSortExpressionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSortExpressionsResponseBody) SetResult(v []*ListSortExpressionsResponseBodyResult) *ListSortExpressionsResponseBody {
	s.Result = v
	return s
}

type ListSortExpressionsResponseBodyResult struct {
	Active      *bool   `json:"active,omitempty" xml:"active,omitempty"`
	Created     *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated     *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListSortExpressionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSortExpressionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponseBodyResult) SetActive(v bool) *ListSortExpressionsResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetCreated(v int32) *ListSortExpressionsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetDescription(v string) *ListSortExpressionsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetName(v string) *ListSortExpressionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetUpdated(v int32) *ListSortExpressionsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListSortExpressionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSortExpressionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSortExpressionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSortExpressionsResponse) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponse) SetHeaders(v map[string]*string) *ListSortExpressionsResponse {
	s.Headers = v
	return s
}

func (s *ListSortExpressionsResponse) SetBody(v *ListSortExpressionsResponseBody) *ListSortExpressionsResponse {
	s.Body = v
	return s
}

type ListSortScriptsResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSortScriptsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSortScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSortScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponseBody) SetRequestId(v string) *ListSortScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSortScriptsResponseBody) SetResult(v []*ListSortScriptsResponseBodyResult) *ListSortScriptsResponseBody {
	s.Result = v
	return s
}

type ListSortScriptsResponseBodyResult struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Scope      *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSortScriptsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSortScriptsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponseBodyResult) SetCreateTime(v string) *ListSortScriptsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetModifyTime(v string) *ListSortScriptsResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetScope(v string) *ListSortScriptsResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetScriptName(v string) *ListSortScriptsResponseBodyResult {
	s.ScriptName = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetStatus(v string) *ListSortScriptsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetType(v string) *ListSortScriptsResponseBodyResult {
	s.Type = &v
	return s
}

type ListSortScriptsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSortScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSortScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSortScriptsResponse) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponse) SetHeaders(v map[string]*string) *ListSortScriptsResponse {
	s.Headers = v
	return s
}

func (s *ListSortScriptsResponse) SetBody(v *ListSortScriptsResponseBody) *ListSortScriptsResponse {
	s.Body = v
	return s
}

type ListStatisticLogsRequest struct {
	Columns    *string `json:"columns,omitempty" xml:"columns,omitempty"`
	Distinct   *bool   `json:"distinct,omitempty" xml:"distinct,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
	SortBy     *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	StartTime  *int32  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StopTime   *int32  `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
}

func (s ListStatisticLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticLogsRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsRequest) SetColumns(v string) *ListStatisticLogsRequest {
	s.Columns = &v
	return s
}

func (s *ListStatisticLogsRequest) SetDistinct(v bool) *ListStatisticLogsRequest {
	s.Distinct = &v
	return s
}

func (s *ListStatisticLogsRequest) SetPageNumber(v int32) *ListStatisticLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStatisticLogsRequest) SetPageSize(v int32) *ListStatisticLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStatisticLogsRequest) SetQuery(v string) *ListStatisticLogsRequest {
	s.Query = &v
	return s
}

func (s *ListStatisticLogsRequest) SetSortBy(v string) *ListStatisticLogsRequest {
	s.SortBy = &v
	return s
}

func (s *ListStatisticLogsRequest) SetStartTime(v int32) *ListStatisticLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListStatisticLogsRequest) SetStopTime(v int32) *ListStatisticLogsRequest {
	s.StopTime = &v
	return s
}

type ListStatisticLogsResponseBody struct {
	RequestId  *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStatisticLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsResponseBody) SetRequestId(v string) *ListStatisticLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticLogsResponseBody) SetResult(v []map[string]interface{}) *ListStatisticLogsResponseBody {
	s.Result = v
	return s
}

func (s *ListStatisticLogsResponseBody) SetTotalCount(v int64) *ListStatisticLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListStatisticLogsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStatisticLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStatisticLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticLogsResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsResponse) SetHeaders(v map[string]*string) *ListStatisticLogsResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticLogsResponse) SetBody(v *ListStatisticLogsResponseBody) *ListStatisticLogsResponse {
	s.Body = v
	return s
}

type ListStatisticReportRequest struct {
	Columns    *string `json:"columns,omitempty" xml:"columns,omitempty"`
	EndTime    *int32  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
	StartTime  *int32  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListStatisticReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticReportRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticReportRequest) SetColumns(v string) *ListStatisticReportRequest {
	s.Columns = &v
	return s
}

func (s *ListStatisticReportRequest) SetEndTime(v int32) *ListStatisticReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListStatisticReportRequest) SetPageNumber(v int32) *ListStatisticReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStatisticReportRequest) SetPageSize(v int32) *ListStatisticReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListStatisticReportRequest) SetQuery(v string) *ListStatisticReportRequest {
	s.Query = &v
	return s
}

func (s *ListStatisticReportRequest) SetStartTime(v int32) *ListStatisticReportRequest {
	s.StartTime = &v
	return s
}

type ListStatisticReportResponseBody struct {
	RequestId  *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStatisticReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticReportResponseBody) SetRequestId(v string) *ListStatisticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticReportResponseBody) SetResult(v []map[string]interface{}) *ListStatisticReportResponseBody {
	s.Result = v
	return s
}

func (s *ListStatisticReportResponseBody) SetTotalCount(v int64) *ListStatisticReportResponseBody {
	s.TotalCount = &v
	return s
}

type ListStatisticReportResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStatisticReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStatisticReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticReportResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticReportResponse) SetHeaders(v map[string]*string) *ListStatisticReportResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticReportResponse) SetBody(v *ListStatisticReportResponseBody) *ListStatisticReportResponse {
	s.Body = v
	return s
}

type ListUserAnalyzerEntriesRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Word       *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListUserAnalyzerEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzerEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesRequest) SetPageNumber(v int32) *ListUserAnalyzerEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) SetPageSize(v int32) *ListUserAnalyzerEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) SetWord(v string) *ListUserAnalyzerEntriesRequest {
	s.Word = &v
	return s
}

type ListUserAnalyzerEntriesResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListUserAnalyzerEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzerEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesResponseBody) SetRequestId(v string) *ListUserAnalyzerEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAnalyzerEntriesResponseBody) SetResult(v map[string]interface{}) *ListUserAnalyzerEntriesResponseBody {
	s.Result = v
	return s
}

type ListUserAnalyzerEntriesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserAnalyzerEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzerEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesResponse) SetHeaders(v map[string]*string) *ListUserAnalyzerEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListUserAnalyzerEntriesResponse) SetBody(v *ListUserAnalyzerEntriesResponseBody) *ListUserAnalyzerEntriesResponse {
	s.Body = v
	return s
}

type ListUserAnalyzersRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUserAnalyzersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersRequest) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersRequest) SetPageNumber(v int32) *ListUserAnalyzersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserAnalyzersRequest) SetPageSize(v int32) *ListUserAnalyzersRequest {
	s.PageSize = &v
	return s
}

type ListUserAnalyzersResponseBody struct {
	RequestId  *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListUserAnalyzersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUserAnalyzersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBody) SetRequestId(v string) *ListUserAnalyzersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAnalyzersResponseBody) SetResult(v []*ListUserAnalyzersResponseBodyResult) *ListUserAnalyzersResponseBody {
	s.Result = v
	return s
}

func (s *ListUserAnalyzersResponseBody) SetTotalCount(v int32) *ListUserAnalyzersResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserAnalyzersResponseBodyResult struct {
	Available *bool                                       `json:"available,omitempty" xml:"available,omitempty"`
	Business  *string                                     `json:"business,omitempty" xml:"business,omitempty"`
	Created   *int32                                      `json:"created,omitempty" xml:"created,omitempty"`
	Dicts     []*ListUserAnalyzersResponseBodyResultDicts `json:"dicts,omitempty" xml:"dicts,omitempty" type:"Repeated"`
	Id        *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	Updated   *int32                                      `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListUserAnalyzersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBodyResult) SetAvailable(v bool) *ListUserAnalyzersResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetBusiness(v string) *ListUserAnalyzersResponseBodyResult {
	s.Business = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetCreated(v int32) *ListUserAnalyzersResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetDicts(v []*ListUserAnalyzersResponseBodyResultDicts) *ListUserAnalyzersResponseBodyResult {
	s.Dicts = v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetId(v string) *ListUserAnalyzersResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetName(v string) *ListUserAnalyzersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetUpdated(v int32) *ListUserAnalyzersResponseBodyResult {
	s.Updated = &v
	return s
}

type ListUserAnalyzersResponseBodyResultDicts struct {
	Available    *bool   `json:"available,omitempty" xml:"available,omitempty"`
	Created      *int32  `json:"created,omitempty" xml:"created,omitempty"`
	EntriesCount *int32  `json:"entriesCount,omitempty" xml:"entriesCount,omitempty"`
	EntriesLimit *int32  `json:"entriesLimit,omitempty" xml:"entriesLimit,omitempty"`
	Id           *string `json:"id,omitempty" xml:"id,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated      *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListUserAnalyzersResponseBodyResultDicts) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponseBodyResultDicts) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetAvailable(v bool) *ListUserAnalyzersResponseBodyResultDicts {
	s.Available = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetCreated(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.Created = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetEntriesCount(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.EntriesCount = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetEntriesLimit(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.EntriesLimit = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetId(v string) *ListUserAnalyzersResponseBodyResultDicts {
	s.Id = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetType(v string) *ListUserAnalyzersResponseBodyResultDicts {
	s.Type = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetUpdated(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.Updated = &v
	return s
}

type ListUserAnalyzersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserAnalyzersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserAnalyzersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponse) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponse) SetHeaders(v map[string]*string) *ListUserAnalyzersResponse {
	s.Headers = v
	return s
}

func (s *ListUserAnalyzersResponse) SetBody(v *ListUserAnalyzersResponseBody) *ListUserAnalyzersResponse {
	s.Body = v
	return s
}

type ModifyAppGroupResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ModifyAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBody) SetRequestId(v string) *ModifyAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppGroupResponseBody) SetResult(v *ModifyAppGroupResponseBodyResult) *ModifyAppGroupResponseBody {
	s.Result = v
	return s
}

type ModifyAppGroupResponseBodyResult struct {
	ChargeType                        *string                                `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay                       *int32                                 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	CommodityCode                     *string                                `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Created                           *int32                                 `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion                    *string                                `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Description                       *string                                `json:"description,omitempty" xml:"description,omitempty"`
	Domain                            *string                                `json:"domain,omitempty" xml:"domain,omitempty"`
	ExpireOn                          *string                                `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	FirstRankAlgoDeploymentId         *int32                                 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	HasPendingQuotaReviewTask         *int32                                 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	Id                                *string                                `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId                        *string                                `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockMode                          *string                                `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration                *int32                                 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	Name                              *string                                `json:"name,omitempty" xml:"name,omitempty"`
	PendingSecondRankAlgoDeploymentId *int32                                 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	ProcessingOrderId                 *string                                `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	Produced                          *int32                                 `json:"produced,omitempty" xml:"produced,omitempty"`
	ProjectId                         *string                                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Quota                             *ModifyAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	SecondRankAlgoDeploymentId        *int32                                 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	Status                            *string                                `json:"status,omitempty" xml:"status,omitempty"`
	SwitchedTime                      *int32                                 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Type                              *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Updated                           *int32                                 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBodyResult) SetChargeType(v string) *ModifyAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetChargingWay(v int32) *ModifyAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCommodityCode(v string) *ModifyAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCreated(v int32) *ModifyAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCurrentVersion(v string) *ModifyAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetDescription(v string) *ModifyAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetDomain(v string) *ModifyAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ModifyAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ModifyAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetId(v string) *ModifyAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetInstanceId(v string) *ModifyAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetLockMode(v string) *ModifyAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *ModifyAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetName(v string) *ModifyAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProcessingOrderId(v string) *ModifyAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProduced(v int32) *ModifyAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProjectId(v string) *ModifyAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetQuota(v *ModifyAppGroupResponseBodyResultQuota) *ModifyAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetStatus(v string) *ModifyAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetSwitchedTime(v int32) *ModifyAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetType(v string) *ModifyAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetUpdated(v int32) *ModifyAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifyAppGroupResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ModifyAppGroupResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *ModifyAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetDocSize(v int32) *ModifyAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetSpec(v string) *ModifyAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ModifyAppGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponse) SetHeaders(v map[string]*string) *ModifyAppGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppGroupResponse) SetBody(v *ModifyAppGroupResponseBody) *ModifyAppGroupResponse {
	s.Body = v
	return s
}

type ModifyAppGroupQuotaResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ModifyAppGroupQuotaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyAppGroupQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBody) SetRequestId(v string) *ModifyAppGroupQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBody) SetResult(v *ModifyAppGroupQuotaResponseBodyResult) *ModifyAppGroupQuotaResponseBody {
	s.Result = v
	return s
}

type ModifyAppGroupQuotaResponseBodyResult struct {
	ChargeType                        *string                                     `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay                       *int32                                      `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	CommodityCode                     *string                                     `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Created                           *int32                                      `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion                    *string                                     `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Description                       *string                                     `json:"description,omitempty" xml:"description,omitempty"`
	ExpireOn                          *string                                     `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	FirstRankAlgoDeploymentId         *int32                                      `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	HasPendingQuotaReviewTask         *int32                                      `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	Id                                *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId                        *string                                     `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockMode                          *string                                     `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration                *int32                                      `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	Name                              *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	PendingSecondRankAlgoDeploymentId *int32                                      `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	ProcessingOrderId                 *string                                     `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	Produced                          *int32                                      `json:"produced,omitempty" xml:"produced,omitempty"`
	ProjectId                         *string                                     `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Quota                             *ModifyAppGroupQuotaResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	SecondRankAlgoDeploymentId        *int32                                      `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	Status                            *string                                     `json:"status,omitempty" xml:"status,omitempty"`
	SwitchedTime                      *int32                                      `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Type                              *string                                     `json:"type,omitempty" xml:"type,omitempty"`
	Updated                           *int32                                      `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyAppGroupQuotaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetChargeType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetChargingWay(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCommodityCode(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCreated(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCurrentVersion(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetDescription(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetInstanceId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetLockMode(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetLockedByExpiration(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetName(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProcessingOrderId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProduced(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProjectId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetQuota(v *ModifyAppGroupQuotaResponseBodyResultQuota) *ModifyAppGroupQuotaResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetStatus(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetSwitchedTime(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetUpdated(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifyAppGroupQuotaResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ModifyAppGroupQuotaResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetComputeResource(v int32) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetDocSize(v int32) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetSpec(v string) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ModifyAppGroupQuotaResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAppGroupQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppGroupQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponse) SetHeaders(v map[string]*string) *ModifyAppGroupQuotaResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppGroupQuotaResponse) SetBody(v *ModifyAppGroupQuotaResponseBody) *ModifyAppGroupQuotaResponse {
	s.Body = v
	return s
}

type ModifyFirstRankRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyFirstRankRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankRequest) SetDryRun(v bool) *ModifyFirstRankRequest {
	s.DryRun = &v
	return s
}

type ModifyFirstRankResponseBody struct {
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ModifyFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBody) SetRequestId(v string) *ModifyFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFirstRankResponseBody) SetResult(v *ModifyFirstRankResponseBodyResult) *ModifyFirstRankResponseBody {
	s.Result = v
	return s
}

type ModifyFirstRankResponseBodyResult struct {
	Active      *bool                                    `json:"active,omitempty" xml:"active,omitempty"`
	Description *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	Meta        []*ModifyFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name        *string                                  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBodyResult) SetActive(v bool) *ModifyFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetDescription(v string) *ModifyFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetMeta(v []*ModifyFirstRankResponseBodyResultMeta) *ModifyFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetName(v string) *ModifyFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type ModifyFirstRankResponseBodyResultMeta struct {
	Arg       *string  `json:"arg,omitempty" xml:"arg,omitempty"`
	Attribute *string  `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Weight    *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ModifyFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetArg(v string) *ModifyFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetAttribute(v string) *ModifyFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetWeight(v float32) *ModifyFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type ModifyFirstRankResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponse) SetHeaders(v map[string]*string) *ModifyFirstRankResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirstRankResponse) SetBody(v *ModifyFirstRankResponseBody) *ModifyFirstRankResponse {
	s.Body = v
	return s
}

type ModifyModelResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelResponseBody) SetRequestId(v string) *ModifyModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyModelResponseBody) SetResult(v string) *ModifyModelResponseBody {
	s.Result = &v
	return s
}

type ModifyModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyModelResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyModelResponse) GoString() string {
	return s.String()
}

func (s *ModifyModelResponse) SetHeaders(v map[string]*string) *ModifyModelResponse {
	s.Headers = v
	return s
}

func (s *ModifyModelResponse) SetBody(v *ModifyModelResponseBody) *ModifyModelResponse {
	s.Body = v
	return s
}

type ModifyQueryProcessorRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyQueryProcessorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorRequest) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorRequest) SetDryRun(v bool) *ModifyQueryProcessorRequest {
	s.DryRun = &v
	return s
}

type ModifyQueryProcessorResponseBody struct {
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ModifyQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponseBody) SetRequestId(v string) *ModifyQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQueryProcessorResponseBody) SetResult(v *ModifyQueryProcessorResponseBodyResult) *ModifyQueryProcessorResponseBody {
	s.Result = v
	return s
}

type ModifyQueryProcessorResponseBodyResult struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Created    *int32                   `json:"created,omitempty" xml:"created,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	Updated    *int32                   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyQueryProcessorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponseBodyResult) SetActive(v bool) *ModifyQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetCreated(v int32) *ModifyQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetDomain(v string) *ModifyQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetIndexes(v []*string) *ModifyQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetName(v string) *ModifyQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *ModifyQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetUpdated(v int32) *ModifyQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifyQueryProcessorResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponse) SetHeaders(v map[string]*string) *ModifyQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *ModifyQueryProcessorResponse) SetBody(v *ModifyQueryProcessorResponseBody) *ModifyQueryProcessorResponse {
	s.Body = v
	return s
}

type ModifyScheduledTaskResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetResult(v map[string]interface{}) *ModifyScheduledTaskResponseBody {
	s.Result = v
	return s
}

type ModifyScheduledTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponse) SetHeaders(v map[string]*string) *ModifyScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledTaskResponse) SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse {
	s.Body = v
	return s
}

type ModifySecondRankRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifySecondRankRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankRequest) GoString() string {
	return s.String()
}

func (s *ModifySecondRankRequest) SetDryRun(v bool) *ModifySecondRankRequest {
	s.DryRun = &v
	return s
}

type ModifySecondRankResponseBody struct {
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ModifySecondRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifySecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponseBody) SetRequestId(v string) *ModifySecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecondRankResponseBody) SetResult(v *ModifySecondRankResponseBodyResult) *ModifySecondRankResponseBody {
	s.Result = v
	return s
}

type ModifySecondRankResponseBodyResult struct {
	Active      *bool   `json:"active,omitempty" xml:"active,omitempty"`
	Created     *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDefault   *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	IsSys       *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	Meta        *string `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated     *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifySecondRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponseBodyResult) SetActive(v bool) *ModifySecondRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetCreated(v int32) *ModifySecondRankResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetDescription(v string) *ModifySecondRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetId(v string) *ModifySecondRankResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetIsDefault(v string) *ModifySecondRankResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetIsSys(v string) *ModifySecondRankResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetMeta(v string) *ModifySecondRankResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetName(v string) *ModifySecondRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetUpdated(v int32) *ModifySecondRankResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifySecondRankResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankResponse) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponse) SetHeaders(v map[string]*string) *ModifySecondRankResponse {
	s.Headers = v
	return s
}

func (s *ModifySecondRankResponse) SetBody(v *ModifySecondRankResponseBody) *ModifySecondRankResponse {
	s.Body = v
	return s
}

type PreviewModelRequest struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s PreviewModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewModelRequest) GoString() string {
	return s.String()
}

func (s *PreviewModelRequest) SetQuery(v string) *PreviewModelRequest {
	s.Query = &v
	return s
}

type PreviewModelResponseBody struct {
	RequestId  *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PreviewModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewModelResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewModelResponseBody) SetRequestId(v string) *PreviewModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewModelResponseBody) SetResult(v []map[string]interface{}) *PreviewModelResponseBody {
	s.Result = v
	return s
}

func (s *PreviewModelResponseBody) SetTotalCount(v int64) *PreviewModelResponseBody {
	s.TotalCount = &v
	return s
}

type PreviewModelResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PreviewModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviewModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewModelResponse) GoString() string {
	return s.String()
}

func (s *PreviewModelResponse) SetHeaders(v map[string]*string) *PreviewModelResponse {
	s.Headers = v
	return s
}

func (s *PreviewModelResponse) SetBody(v *PreviewModelResponseBody) *PreviewModelResponse {
	s.Body = v
	return s
}

type PushInterventionDictionaryEntriesResponseBody struct {
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s PushInterventionDictionaryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionDictionaryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesResponseBody) SetRequestId(v string) *PushInterventionDictionaryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushInterventionDictionaryEntriesResponseBody) SetResult(v []*string) *PushInterventionDictionaryEntriesResponseBody {
	s.Result = v
	return s
}

type PushInterventionDictionaryEntriesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushInterventionDictionaryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionDictionaryEntriesResponse) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesResponse) SetHeaders(v map[string]*string) *PushInterventionDictionaryEntriesResponse {
	s.Headers = v
	return s
}

func (s *PushInterventionDictionaryEntriesResponse) SetBody(v *PushInterventionDictionaryEntriesResponseBody) *PushInterventionDictionaryEntriesResponse {
	s.Body = v
	return s
}

type PushUserAnalyzerEntriesResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushUserAnalyzerEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushUserAnalyzerEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesResponseBody) SetRequestId(v string) *PushUserAnalyzerEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushUserAnalyzerEntriesResponseBody) SetResult(v map[string]interface{}) *PushUserAnalyzerEntriesResponseBody {
	s.Result = v
	return s
}

type PushUserAnalyzerEntriesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushUserAnalyzerEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s PushUserAnalyzerEntriesResponse) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesResponse) SetHeaders(v map[string]*string) *PushUserAnalyzerEntriesResponse {
	s.Headers = v
	return s
}

func (s *PushUserAnalyzerEntriesResponse) SetBody(v *PushUserAnalyzerEntriesResponseBody) *PushUserAnalyzerEntriesResponse {
	s.Body = v
	return s
}

type RankPreviewQueryResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RankPreviewQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RankPreviewQueryResponseBody) GoString() string {
	return s.String()
}

func (s *RankPreviewQueryResponseBody) SetRequestId(v string) *RankPreviewQueryResponseBody {
	s.RequestId = &v
	return s
}

type RankPreviewQueryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RankPreviewQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RankPreviewQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s RankPreviewQueryResponse) GoString() string {
	return s.String()
}

func (s *RankPreviewQueryResponse) SetHeaders(v map[string]*string) *RankPreviewQueryResponse {
	s.Headers = v
	return s
}

func (s *RankPreviewQueryResponse) SetBody(v *RankPreviewQueryResponseBody) *RankPreviewQueryResponse {
	s.Body = v
	return s
}

type ReleaseSortScriptResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReleaseSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponseBody) SetRequestId(v string) *ReleaseSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseSortScriptResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSortScriptResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponse) SetHeaders(v map[string]*string) *ReleaseSortScriptResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSortScriptResponse) SetBody(v *ReleaseSortScriptResponseBody) *ReleaseSortScriptResponse {
	s.Body = v
	return s
}

type RemoveAppResponseBody struct {
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppResponseBody) SetRequestId(v string) *RemoveAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppResponseBody) SetResult(v []*int32) *RemoveAppResponseBody {
	s.Result = v
	return s
}

type RemoveAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAppResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppResponse) SetHeaders(v map[string]*string) *RemoveAppResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppResponse) SetBody(v *RemoveAppResponseBody) *RemoveAppResponse {
	s.Body = v
	return s
}

type RemoveAppGroupResponseBody struct {
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppGroupResponseBody) SetRequestId(v string) *RemoveAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppGroupResponseBody) SetResult(v []*int32) *RemoveAppGroupResponseBody {
	s.Result = v
	return s
}

type RemoveAppGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppGroupResponse) SetHeaders(v map[string]*string) *RemoveAppGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppGroupResponse) SetBody(v *RemoveAppGroupResponseBody) *RemoveAppGroupResponse {
	s.Body = v
	return s
}

type RemoveDataCollectionResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveDataCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDataCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataCollectionResponseBody) SetRequestId(v string) *RemoveDataCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataCollectionResponseBody) SetResult(v string) *RemoveDataCollectionResponseBody {
	s.Result = &v
	return s
}

type RemoveDataCollectionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDataCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDataCollectionResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataCollectionResponse) SetHeaders(v map[string]*string) *RemoveDataCollectionResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataCollectionResponse) SetBody(v *RemoveDataCollectionResponseBody) *RemoveDataCollectionResponse {
	s.Body = v
	return s
}

type RemoveFirstRankResponseBody struct {
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *RemoveFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RemoveFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBody) SetRequestId(v string) *RemoveFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFirstRankResponseBody) SetResult(v *RemoveFirstRankResponseBodyResult) *RemoveFirstRankResponseBody {
	s.Result = v
	return s
}

type RemoveFirstRankResponseBodyResult struct {
	Active      *bool                                    `json:"active,omitempty" xml:"active,omitempty"`
	Description *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	Meta        []*RemoveFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name        *string                                  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RemoveFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBodyResult) SetActive(v bool) *RemoveFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetDescription(v string) *RemoveFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetMeta(v []*RemoveFirstRankResponseBodyResultMeta) *RemoveFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetName(v string) *RemoveFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type RemoveFirstRankResponseBodyResultMeta struct {
	Arg       *string  `json:"arg,omitempty" xml:"arg,omitempty"`
	Attribute *string  `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Weight    *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s RemoveFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetArg(v string) *RemoveFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetAttribute(v string) *RemoveFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetWeight(v float32) *RemoveFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type RemoveFirstRankResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponse) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponse) SetHeaders(v map[string]*string) *RemoveFirstRankResponse {
	s.Headers = v
	return s
}

func (s *RemoveFirstRankResponse) SetBody(v *RemoveFirstRankResponseBody) *RemoveFirstRankResponse {
	s.Body = v
	return s
}

type RemoveInterventionDictionaryResponseBody struct {
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *RemoveInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RemoveInterventionDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponseBody) SetRequestId(v string) *RemoveInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBody) SetResult(v *RemoveInterventionDictionaryResponseBodyResult) *RemoveInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

type RemoveInterventionDictionaryResponseBodyResult struct {
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	Created  *string `json:"created,omitempty" xml:"created,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated  *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s RemoveInterventionDictionaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetCreated(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetName(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetType(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetUpdated(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

type RemoveInterventionDictionaryResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveInterventionDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponse) SetHeaders(v map[string]*string) *RemoveInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *RemoveInterventionDictionaryResponse) SetBody(v *RemoveInterventionDictionaryResponseBody) *RemoveInterventionDictionaryResponse {
	s.Body = v
	return s
}

type RemoveQueryProcessorResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveQueryProcessorResponseBody) SetRequestId(v string) *RemoveQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveQueryProcessorResponseBody) SetResult(v string) *RemoveQueryProcessorResponseBody {
	s.Result = &v
	return s
}

type RemoveQueryProcessorResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *RemoveQueryProcessorResponse) SetHeaders(v map[string]*string) *RemoveQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *RemoveQueryProcessorResponse) SetBody(v *RemoveQueryProcessorResponseBody) *RemoveQueryProcessorResponse {
	s.Body = v
	return s
}

type RemoveScheduledTaskResponseBody struct {
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveScheduledTaskResponseBody) SetRequestId(v string) *RemoveScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveScheduledTaskResponseBody) SetResult(v []*int32) *RemoveScheduledTaskResponseBody {
	s.Result = v
	return s
}

type RemoveScheduledTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *RemoveScheduledTaskResponse) SetHeaders(v map[string]*string) *RemoveScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *RemoveScheduledTaskResponse) SetBody(v *RemoveScheduledTaskResponseBody) *RemoveScheduledTaskResponse {
	s.Body = v
	return s
}

type RemoveSearchStrategyResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSearchStrategyResponseBody) SetRequestId(v string) *RemoveSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSearchStrategyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *RemoveSearchStrategyResponse) SetHeaders(v map[string]*string) *RemoveSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *RemoveSearchStrategyResponse) SetBody(v *RemoveSearchStrategyResponseBody) *RemoveSearchStrategyResponse {
	s.Body = v
	return s
}

type RemoveSecondRankResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveSecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSecondRankResponseBody) SetRequestId(v string) *RemoveSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSecondRankResponseBody) SetResult(v map[string]interface{}) *RemoveSecondRankResponseBody {
	s.Result = v
	return s
}

type RemoveSecondRankResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSecondRankResponse) GoString() string {
	return s.String()
}

func (s *RemoveSecondRankResponse) SetHeaders(v map[string]*string) *RemoveSecondRankResponse {
	s.Headers = v
	return s
}

func (s *RemoveSecondRankResponse) SetBody(v *RemoveSecondRankResponseBody) *RemoveSecondRankResponse {
	s.Body = v
	return s
}

type RemoveUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserAnalyzerResponseBody) SetRequestId(v string) *RemoveUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *RemoveUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type RemoveUserAnalyzerResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserAnalyzerResponse) SetHeaders(v map[string]*string) *RemoveUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserAnalyzerResponse) SetBody(v *RemoveUserAnalyzerResponseBody) *RemoveUserAnalyzerResponse {
	s.Body = v
	return s
}

type RenewAppGroupResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RenewAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppGroupResponseBody) SetRequestId(v string) *RenewAppGroupResponseBody {
	s.RequestId = &v
	return s
}

type RenewAppGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewAppGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewAppGroupResponse) SetHeaders(v map[string]*string) *RenewAppGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewAppGroupResponse) SetBody(v *RenewAppGroupResponseBody) *RenewAppGroupResponse {
	s.Body = v
	return s
}

type ReplaceAppGroupCommodityCodeResponseBody struct {
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ReplaceAppGroupCommodityCodeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ReplaceAppGroupCommodityCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) SetRequestId(v string) *ReplaceAppGroupCommodityCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) SetResult(v *ReplaceAppGroupCommodityCodeResponseBodyResult) *ReplaceAppGroupCommodityCodeResponseBody {
	s.Result = v
	return s
}

type ReplaceAppGroupCommodityCodeResponseBodyResult struct {
	ChargeType                        *string                                              `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay                       *int32                                               `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	CommodityCode                     *string                                              `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Created                           *int32                                               `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion                    *string                                              `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Description                       *string                                              `json:"description,omitempty" xml:"description,omitempty"`
	ExpireOn                          *string                                              `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	FirstRankAlgoDeploymentId         *int32                                               `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	HasPendingQuotaReviewTask         *int32                                               `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	Id                                *string                                              `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId                        *string                                              `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockMode                          *string                                              `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration                *int32                                               `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	Name                              *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	PendingSecondRankAlgoDeploymentId *int32                                               `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	ProcessingOrderId                 *string                                              `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	Produced                          *int32                                               `json:"produced,omitempty" xml:"produced,omitempty"`
	ProjectId                         *string                                              `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Quota                             *ReplaceAppGroupCommodityCodeResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	SecondRankAlgoDeploymentId        *int32                                               `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	Status                            *string                                              `json:"status,omitempty" xml:"status,omitempty"`
	SwitchedTime                      *int32                                               `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Type                              *string                                              `json:"type,omitempty" xml:"type,omitempty"`
	Updated                           *int32                                               `json:"updated,omitempty" xml:"updated,omitempty"`
	Versions                          []*string                                            `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetChargeType(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetChargingWay(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCommodityCode(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCreated(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCurrentVersion(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetDescription(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetExpireOn(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetInstanceId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetLockMode(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetLockedByExpiration(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetName(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProcessingOrderId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProduced(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProjectId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetQuota(v *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetStatus(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetSwitchedTime(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetType(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetUpdated(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetVersions(v []*string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Versions = v
	return s
}

type ReplaceAppGroupCommodityCodeResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetComputeResource(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetDocSize(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetSpec(v string) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ReplaceAppGroupCommodityCodeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplaceAppGroupCommodityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceAppGroupCommodityCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponse) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetHeaders(v map[string]*string) *ReplaceAppGroupCommodityCodeResponse {
	s.Headers = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetBody(v *ReplaceAppGroupCommodityCodeResponseBody) *ReplaceAppGroupCommodityCodeResponse {
	s.Body = v
	return s
}

type SaveSortScriptFileResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SaveSortScriptFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileResponseBody) SetRequestId(v string) *SaveSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

type SaveSortScriptFileResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSortScriptFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileResponse) SetHeaders(v map[string]*string) *SaveSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *SaveSortScriptFileResponse) SetBody(v *SaveSortScriptFileResponseBody) *SaveSortScriptFileResponse {
	s.Body = v
	return s
}

type StartSlowQueryAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartSlowQueryAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSlowQueryAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *StartSlowQueryAnalyzerResponseBody) SetRequestId(v string) *StartSlowQueryAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSlowQueryAnalyzerResponseBody) SetResult(v map[string]interface{}) *StartSlowQueryAnalyzerResponseBody {
	s.Result = v
	return s
}

type StartSlowQueryAnalyzerResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartSlowQueryAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartSlowQueryAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSlowQueryAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *StartSlowQueryAnalyzerResponse) SetHeaders(v map[string]*string) *StartSlowQueryAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *StartSlowQueryAnalyzerResponse) SetBody(v *StartSlowQueryAnalyzerResponseBody) *StartSlowQueryAnalyzerResponse {
	s.Body = v
	return s
}

type TrainModelResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TrainModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrainModelResponseBody) GoString() string {
	return s.String()
}

func (s *TrainModelResponseBody) SetRequestId(v string) *TrainModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainModelResponseBody) SetResult(v map[string]interface{}) *TrainModelResponseBody {
	s.Result = v
	return s
}

type TrainModelResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TrainModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrainModelResponse) String() string {
	return tea.Prettify(s)
}

func (s TrainModelResponse) GoString() string {
	return s.String()
}

func (s *TrainModelResponse) SetHeaders(v map[string]*string) *TrainModelResponse {
	s.Headers = v
	return s
}

func (s *TrainModelResponse) SetBody(v *TrainModelResponseBody) *TrainModelResponse {
	s.Body = v
	return s
}

type UnbindESUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnbindESUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindESUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerResponseBody) SetRequestId(v string) *UnbindESUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindESUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *UnbindESUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type UnbindESUserAnalyzerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindESUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindESUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerResponse) SetHeaders(v map[string]*string) *UnbindESUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *UnbindESUserAnalyzerResponse) SetBody(v *UnbindESUserAnalyzerResponseBody) *UnbindESUserAnalyzerResponse {
	s.Body = v
	return s
}

type UnbindEsInstanceResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnbindEsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEsInstanceResponseBody) SetRequestId(v string) *UnbindEsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEsInstanceResponseBody) SetResult(v map[string]interface{}) *UnbindEsInstanceResponseBody {
	s.Result = v
	return s
}

type UnbindEsInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindEsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEsInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnbindEsInstanceResponse) SetHeaders(v map[string]*string) *UnbindEsInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnbindEsInstanceResponse) SetBody(v *UnbindEsInstanceResponseBody) *UnbindEsInstanceResponse {
	s.Body = v
	return s
}

type UpdateABTestExperimentResponseBody struct {
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponseBody) SetRequestId(v string) *UpdateABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestExperimentResponseBody) SetResult(v *UpdateABTestExperimentResponseBodyResult) *UpdateABTestExperimentResponseBody {
	s.Result = v
	return s
}

type UpdateABTestExperimentResponseBodyResult struct {
	Created *int32                 `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string                `json:"name,omitempty" xml:"name,omitempty"`
	Online  *bool                  `json:"online,omitempty" xml:"online,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Traffic *int32                 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	Updated *int32                 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestExperimentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponseBodyResult) SetCreated(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetId(v string) *UpdateABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetName(v string) *UpdateABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetOnline(v bool) *UpdateABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetParams(v map[string]interface{}) *UpdateABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetTraffic(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetUpdated(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateABTestExperimentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponse) SetHeaders(v map[string]*string) *UpdateABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestExperimentResponse) SetBody(v *UpdateABTestExperimentResponseBody) *UpdateABTestExperimentResponse {
	s.Body = v
	return s
}

type UpdateABTestFixedFlowDividersResponseBody struct {
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s UpdateABTestFixedFlowDividersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersResponseBody) SetRequestId(v string) *UpdateABTestFixedFlowDividersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponseBody) SetResult(v []*string) *UpdateABTestFixedFlowDividersResponseBody {
	s.Result = v
	return s
}

type UpdateABTestFixedFlowDividersResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestFixedFlowDividersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersResponse) SetHeaders(v map[string]*string) *UpdateABTestFixedFlowDividersResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponse) SetBody(v *UpdateABTestFixedFlowDividersResponseBody) *UpdateABTestFixedFlowDividersResponse {
	s.Body = v
	return s
}

type UpdateABTestGroupResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponseBody) SetRequestId(v string) *UpdateABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestGroupResponseBody) SetResult(v *UpdateABTestGroupResponseBodyResult) *UpdateABTestGroupResponseBody {
	s.Result = v
	return s
}

type UpdateABTestGroupResponseBodyResult struct {
	Created *int32  `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Status  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Updated *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponseBodyResult) SetCreated(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetId(v string) *UpdateABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetName(v string) *UpdateABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetStatus(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetUpdated(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateABTestGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponse) SetHeaders(v map[string]*string) *UpdateABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestGroupResponse) SetBody(v *UpdateABTestGroupResponseBody) *UpdateABTestGroupResponse {
	s.Body = v
	return s
}

type UpdateABTestSceneResponseBody struct {
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponseBody) SetRequestId(v string) *UpdateABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestSceneResponseBody) SetResult(v *UpdateABTestSceneResponseBodyResult) *UpdateABTestSceneResponseBody {
	s.Result = v
	return s
}

type UpdateABTestSceneResponseBodyResult struct {
	Created *int32                 `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string                `json:"name,omitempty" xml:"name,omitempty"`
	Online  *bool                  `json:"online,omitempty" xml:"online,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Traffic *int32                 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	Updated *int32                 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponseBodyResult) SetCreated(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetId(v string) *UpdateABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetName(v string) *UpdateABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetOnline(v bool) *UpdateABTestSceneResponseBodyResult {
	s.Online = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetParams(v map[string]interface{}) *UpdateABTestSceneResponseBodyResult {
	s.Params = v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetTraffic(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetUpdated(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateABTestSceneResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponse) SetHeaders(v map[string]*string) *UpdateABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestSceneResponse) SetBody(v *UpdateABTestSceneResponseBody) *UpdateABTestSceneResponse {
	s.Body = v
	return s
}

type UpdateFetchFieldsRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateFetchFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFetchFieldsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsRequest) SetDryRun(v bool) *UpdateFetchFieldsRequest {
	s.DryRun = &v
	return s
}

type UpdateFetchFieldsResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateFetchFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFetchFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsResponseBody) SetRequestId(v string) *UpdateFetchFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFetchFieldsResponseBody) SetResult(v bool) *UpdateFetchFieldsResponseBody {
	s.Result = &v
	return s
}

type UpdateFetchFieldsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFetchFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFetchFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFetchFieldsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsResponse) SetHeaders(v map[string]*string) *UpdateFetchFieldsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFetchFieldsResponse) SetBody(v *UpdateFetchFieldsResponseBody) *UpdateFetchFieldsResponse {
	s.Body = v
	return s
}

type UpdateFunctionDefaultInstanceRequest struct {
	// 实例名称
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s UpdateFunctionDefaultInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceRequest) SetInstanceName(v string) *UpdateFunctionDefaultInstanceRequest {
	s.InstanceName = &v
	return s
}

type UpdateFunctionDefaultInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int64  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Latency   *int64  `json:"Latency,omitempty" xml:"Latency,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionDefaultInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetCode(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetHttpCode(v int64) *UpdateFunctionDefaultInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetLatency(v int64) *UpdateFunctionDefaultInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetMessage(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetRequestId(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetStatus(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Status = &v
	return s
}

type UpdateFunctionDefaultInstanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionDefaultInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceResponse) SetHeaders(v map[string]*string) *UpdateFunctionDefaultInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponse) SetBody(v *UpdateFunctionDefaultInstanceResponseBody) *UpdateFunctionDefaultInstanceResponse {
	s.Body = v
	return s
}

type UpdateFunctionInstanceRequest struct {
	// 创建参数
	CreateParameters []*UpdateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// 周期训练
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// 实例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateFunctionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequest) SetCreateParameters(v []*UpdateFunctionInstanceRequestCreateParameters) *UpdateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetCron(v string) *UpdateFunctionInstanceRequest {
	s.Cron = &v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetDescription(v string) *UpdateFunctionInstanceRequest {
	s.Description = &v
	return s
}

type UpdateFunctionInstanceRequestCreateParameters struct {
	// 参数名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 参数值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateFunctionInstanceRequestCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetName(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetValue(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

type UpdateFunctionInstanceResponseBody struct {
	// 错误码
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode *int64  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 耗时
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// 错误信息
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceResponseBody) SetCode(v string) *UpdateFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetHttpCode(v int64) *UpdateFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetLatency(v int64) *UpdateFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetMessage(v string) *UpdateFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetRequestId(v string) *UpdateFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetStatus(v string) *UpdateFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type UpdateFunctionInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceResponse) SetHeaders(v map[string]*string) *UpdateFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionInstanceResponse) SetBody(v *UpdateFunctionInstanceResponseBody) *UpdateFunctionInstanceResponse {
	s.Body = v
	return s
}

type UpdateSearchStrategyResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyResponseBody) SetRequestId(v string) *UpdateSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSearchStrategyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyResponse) SetHeaders(v map[string]*string) *UpdateSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSearchStrategyResponse) SetBody(v *UpdateSearchStrategyResponseBody) *UpdateSearchStrategyResponse {
	s.Body = v
	return s
}

type UpdateSortScriptResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptResponseBody) SetRequestId(v string) *UpdateSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSortScriptResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSortScriptResponse) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptResponse) SetHeaders(v map[string]*string) *UpdateSortScriptResponse {
	s.Headers = v
	return s
}

func (s *UpdateSortScriptResponse) SetBody(v *UpdateSortScriptResponseBody) *UpdateSortScriptResponse {
	s.Body = v
	return s
}

type UpdateSummariesRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateSummariesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesRequest) GoString() string {
	return s.String()
}

func (s *UpdateSummariesRequest) SetDryRun(v bool) *UpdateSummariesRequest {
	s.DryRun = &v
	return s
}

type UpdateSummariesResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSummariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSummariesResponseBody) SetRequestId(v string) *UpdateSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSummariesResponseBody) SetResult(v bool) *UpdateSummariesResponseBody {
	s.Result = &v
	return s
}

type UpdateSummariesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSummariesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesResponse) GoString() string {
	return s.String()
}

func (s *UpdateSummariesResponse) SetHeaders(v map[string]*string) *UpdateSummariesResponse {
	s.Headers = v
	return s
}

func (s *UpdateSummariesResponse) SetBody(v *UpdateSummariesResponseBody) *UpdateSummariesResponse {
	s.Body = v
	return s
}

type ValidateDataSourcesResponseBody struct {
	// Id of the request
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ValidateDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ValidateDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBody) SetRequestId(v string) *ValidateDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateDataSourcesResponseBody) SetResult(v []*ValidateDataSourcesResponseBodyResult) *ValidateDataSourcesResponseBody {
	s.Result = v
	return s
}

type ValidateDataSourcesResponseBodyResult struct {
	Code       *string                                          `json:"code,omitempty" xml:"code,omitempty"`
	DataSource *ValidateDataSourcesResponseBodyResultDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	Message    *string                                          `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ValidateDataSourcesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBodyResult) SetCode(v string) *ValidateDataSourcesResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) SetDataSource(v *ValidateDataSourcesResponseBodyResultDataSource) *ValidateDataSourcesResponseBodyResult {
	s.DataSource = v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) SetMessage(v string) *ValidateDataSourcesResponseBodyResult {
	s.Message = &v
	return s
}

type ValidateDataSourcesResponseBodyResultDataSource struct {
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	TableName  *string                `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ValidateDataSourcesResponseBodyResultDataSource) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponseBodyResultDataSource) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetParameters(v map[string]interface{}) *ValidateDataSourcesResponseBodyResultDataSource {
	s.Parameters = v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetTableName(v string) *ValidateDataSourcesResponseBodyResultDataSource {
	s.TableName = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetType(v string) *ValidateDataSourcesResponseBodyResultDataSource {
	s.Type = &v
	return s
}

type ValidateDataSourcesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponse) SetHeaders(v map[string]*string) *ValidateDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ValidateDataSourcesResponse) SetBody(v *ValidateDataSourcesResponseBody) *ValidateDataSourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("opensearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindESUserAnalyzer(appGroupIdentity *string, esInstanceId *string) (_result *BindESUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindESUserAnalyzerResponse{}
	_body, _err := client.BindESUserAnalyzerWithOptions(appGroupIdentity, esInstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindESUserAnalyzerWithOptions(appGroupIdentity *string, esInstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindESUserAnalyzerResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	esInstanceId = openapiutil.GetEncodeParam(esInstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("BindESUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/es/" + tea.StringValue(esInstanceId) + "/actions/bind-analyzer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BindESUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindEsInstance(appGroupIdentity *string) (_result *BindEsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindEsInstanceResponse{}
	_body, _err := client.BindEsInstanceWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindEsInstanceWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindEsInstanceResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("BindEsInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/actions/bind-es-instance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompileSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *CompileSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CompileSortScriptResponse{}
	_body, _err := client.CompileSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompileSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CompileSortScriptResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CompileSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName) + "/actions/compiling"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CompileSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string) (_result *CreateABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABTestExperimentResponse{}
	_body, _err := client.CreateABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABTestExperimentResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateABTestGroup(appGroupIdentity *string, sceneId *string) (_result *CreateABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABTestGroupResponse{}
	_body, _err := client.CreateABTestGroupWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABTestGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateABTestScene(appGroupIdentity *string) (_result *CreateABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABTestSceneResponse{}
	_body, _err := client.CreateABTestSceneWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateABTestSceneWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABTestSceneResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(appGroupIdentity *string, request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(appGroupIdentity *string, request *CreateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppGroup() (_result *CreateAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CreateAppGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppGroupWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataCollection(appGroupIdentity *string) (_result *CreateDataCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataCollectionResponse{}
	_body, _err := client.CreateDataCollectionWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataCollectionWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataCollectionResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataCollection"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/data-collections"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFirstRank(appGroupIdentity *string, appId *string, request *CreateFirstRankRequest) (_result *CreateFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFirstRankResponse{}
	_body, _err := client.CreateFirstRankWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFirstRankWithOptions(appGroupIdentity *string, appId *string, request *CreateFirstRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFirstRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/first-ranks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunctionInstance(appGroupIdentity *string, functionName *string, request *CreateFunctionInstanceRequest) (_result *CreateFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CreateFunctionInstanceWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, request *CreateFunctionInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateParameters)) {
		body["createParameters"] = request.CreateParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Cron)) {
		body["cron"] = request.Cron
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		body["functionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		body["modelType"] = request.ModelType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunctionTask(appGroupIdentity *string, functionName *string, instanceName *string) (_result *CreateFunctionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CreateFunctionTaskWithOptions(appGroupIdentity, functionName, instanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionTaskWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionTaskResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	instanceName = openapiutil.GetEncodeParam(instanceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunctionTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances/" + tea.StringValue(instanceName) + "/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInterventionDictionary() (_result *CreateInterventionDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInterventionDictionaryResponse{}
	_body, _err := client.CreateInterventionDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInterventionDictionaryWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInterventionDictionaryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInterventionDictionary"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInterventionDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModel(appGroupIdentity *string) (_result *CreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQueryProcessor(appGroupIdentity *string, appId *string, request *CreateQueryProcessorRequest) (_result *CreateQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQueryProcessorResponse{}
	_body, _err := client.CreateQueryProcessorWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQueryProcessorWithOptions(appGroupIdentity *string, appId *string, request *CreateQueryProcessorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQueryProcessorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/query-processors"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScheduledTask(appGroupIdentity *string) (_result *CreateScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CreateScheduledTaskWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduledTaskWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scheduled-tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSearchStrategy(appGroupIdentity *string, appId *string) (_result *CreateSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSearchStrategyResponse{}
	_body, _err := client.CreateSearchStrategyWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSearchStrategyWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSearchStrategyResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/search-strategies"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecondRank(appGroupIdentity *string, appId *string, request *CreateSecondRankRequest) (_result *CreateSecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSecondRankResponse{}
	_body, _err := client.CreateSecondRankWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecondRankWithOptions(appGroupIdentity *string, appId *string, request *CreateSecondRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSecondRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/second-ranks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSortScript(appGroupIdentity *string, appVersionId *string) (_result *CreateSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSortScriptResponse{}
	_body, _err := client.CreateSortScriptWithOptions(appGroupIdentity, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSortScriptWithOptions(appGroupIdentity *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSortScriptResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserAnalyzer() (_result *CreateUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserAnalyzerResponse{}
	_body, _err := client.CreateUserAnalyzerWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserAnalyzerWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUserAnalyzerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *DeleteABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABTestExperimentResponse{}
	_body, _err := client.DeleteABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABTestExperimentResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	experimentId = openapiutil.GetEncodeParam(experimentId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments/" + tea.StringValue(experimentId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteABTestGroup(appGroupIdentity *string, sceneId *string, groupId *string) (_result *DeleteABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABTestGroupResponse{}
	_body, _err := client.DeleteABTestGroupWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABTestGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteABTestScene(appGroupIdentity *string, sceneId *string) (_result *DeleteABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABTestSceneResponse{}
	_body, _err := client.DeleteABTestSceneWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteABTestSceneWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABTestSceneResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunctionInstance(appGroupIdentity *string, functionName *string, instanceName *string) (_result *DeleteFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.DeleteFunctionInstanceWithOptions(appGroupIdentity, functionName, instanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionInstanceResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	instanceName = openapiutil.GetEncodeParam(instanceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances/" + tea.StringValue(instanceName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModel(appGroupIdentity *string, modelName *string) (_result *DeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *DeleteSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSortScriptResponse{}
	_body, _err := client.DeleteSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSortScriptResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSortScriptFile(appGroupIdentity *string, appVersionId *string, scriptName *string, fileName *string) (_result *DeleteSortScriptFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSortScriptFileResponse{}
	_body, _err := client.DeleteSortScriptFileWithOptions(appGroupIdentity, appVersionId, scriptName, fileName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSortScriptFileWithOptions(appGroupIdentity *string, appVersionId *string, scriptName *string, fileName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSortScriptFileResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	fileName = openapiutil.GetEncodeParam(fileName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSortScriptFile"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName) + "/files/src/" + tea.StringValue(fileName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSortScriptFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *DescribeABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeABTestExperimentResponse{}
	_body, _err := client.DescribeABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeABTestExperimentResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	experimentId = openapiutil.GetEncodeParam(experimentId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments/" + tea.StringValue(experimentId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeABTestGroup(appGroupIdentity *string, sceneId *string, groupId *string) (_result *DescribeABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeABTestGroupResponse{}
	_body, _err := client.DescribeABTestGroupWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeABTestGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeABTestScene(appGroupIdentity *string, sceneId *string) (_result *DescribeABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeABTestSceneResponse{}
	_body, _err := client.DescribeABTestSceneWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeABTestSceneWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeABTestSceneResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApp(appGroupIdentity *string, appId *string) (_result *DescribeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppResponse{}
	_body, _err := client.DescribeAppWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApp"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppGroup(appGroupIdentity *string) (_result *DescribeAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppGroupResponse{}
	_body, _err := client.DescribeAppGroupWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppGroupWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppGroupDataReport(appGroupIdentity *string, request *DescribeAppGroupDataReportRequest) (_result *DescribeAppGroupDataReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppGroupDataReportResponse{}
	_body, _err := client.DescribeAppGroupDataReportWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppGroupDataReportWithOptions(appGroupIdentity *string, request *DescribeAppGroupDataReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppGroupDataReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppGroupDataReport"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/data-report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppGroupDataReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppGroupStatistics(appGroupIdentity *string) (_result *DescribeAppGroupStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppGroupStatisticsResponse{}
	_body, _err := client.DescribeAppGroupStatisticsWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppGroupStatisticsWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppGroupStatisticsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppGroupStatistics"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppGroupStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppStatistics(appGroupIdentity *string, appId *string) (_result *DescribeAppStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppStatisticsResponse{}
	_body, _err := client.DescribeAppStatisticsWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppStatisticsWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppStatisticsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppStatistics"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApps(appGroupIdentity *string) (_result *DescribeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppsWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApps"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataCollction(appGroupIdentity *string, dataCollectionIdentity *string) (_result *DescribeDataCollctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDataCollctionResponse{}
	_body, _err := client.DescribeDataCollctionWithOptions(appGroupIdentity, dataCollectionIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataCollctionWithOptions(appGroupIdentity *string, dataCollectionIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDataCollctionResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	dataCollectionIdentity = openapiutil.GetEncodeParam(dataCollectionIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataCollction"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/data-collections/" + tea.StringValue(dataCollectionIdentity)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataCollctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFirstRank(appGroupIdentity *string, appId *string, name *string) (_result *DescribeFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFirstRankResponse{}
	_body, _err := client.DescribeFirstRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFirstRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFirstRankResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/first-ranks/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInterventionDictionary(name *string) (_result *DescribeInterventionDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInterventionDictionaryResponse{}
	_body, _err := client.DescribeInterventionDictionaryWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInterventionDictionaryWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInterventionDictionaryResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInterventionDictionary"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInterventionDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModel(appGroupIdentity *string, modelName *string) (_result *DescribeModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeModelResponse{}
	_body, _err := client.DescribeModelWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModelWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeModelResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQueryProcessor(appGroupIdentity *string, appId *string, name *string) (_result *DescribeQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQueryProcessorResponse{}
	_body, _err := client.DescribeQueryProcessorWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQueryProcessorWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQueryProcessorResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/query-processors/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegion() (_result *DescribeRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionResponse{}
	_body, _err := client.DescribeRegionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegion"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/region"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduledTask(appGroupIdentity *string, taskId *string) (_result *DescribeScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeScheduledTaskResponse{}
	_body, _err := client.DescribeScheduledTaskWithOptions(appGroupIdentity, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScheduledTaskWithOptions(appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeScheduledTaskResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	taskId = openapiutil.GetEncodeParam(taskId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scheduled-tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecondRank(appGroupIdentity *string, appId *string, name *string) (_result *DescribeSecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSecondRankResponse{}
	_body, _err := client.DescribeSecondRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecondRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSecondRankResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/second-ranks/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowQueryStatus(appGroupIdentity *string) (_result *DescribeSlowQueryStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSlowQueryStatusResponse{}
	_body, _err := client.DescribeSlowQueryStatusWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowQueryStatusWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSlowQueryStatusResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowQueryStatus"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/optimizers/slow-query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowQueryStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserAnalyzer(name *string, request *DescribeUserAnalyzerRequest) (_result *DescribeUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserAnalyzerResponse{}
	_body, _err := client.DescribeUserAnalyzerWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserAnalyzerWithOptions(name *string, request *DescribeUserAnalyzerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUserAnalyzerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.With)) {
		query["with"] = request.With
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSlowQuery(appGroupIdentity *string) (_result *DisableSlowQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableSlowQueryResponse{}
	_body, _err := client.DisableSlowQueryWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSlowQueryWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableSlowQueryResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSlowQuery"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/optimizers/slow-query/actions/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSlowQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSlowQuery(appGroupIdentity *string) (_result *EnableSlowQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableSlowQueryResponse{}
	_body, _err := client.EnableSlowQueryWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSlowQueryWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableSlowQueryResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSlowQuery"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/optimizers/slow-query/actions/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSlowQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateMergedTable(request *GenerateMergedTableRequest) (_result *GenerateMergedTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateMergedTableResponse{}
	_body, _err := client.GenerateMergedTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateMergedTableWithOptions(request *GenerateMergedTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateMergedTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		query["spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateMergedTable"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/schema/actions/merge"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateMergedTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomain(domainName *string, request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(domainName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainWithOptions(domainName *string, request *GetDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	domainName = openapiutil.GetEncodeParam(domainName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppGroupIdentity)) {
		query["appGroupIdentity"] = request.AppGroupIdentity
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomain"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/domains/" + tea.StringValue(domainName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionCurrentVersion(functionName *string, request *GetFunctionCurrentVersionRequest) (_result *GetFunctionCurrentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionCurrentVersionResponse{}
	_body, _err := client.GetFunctionCurrentVersionWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionCurrentVersionWithOptions(functionName *string, request *GetFunctionCurrentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionCurrentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		query["functionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["modelType"] = request.ModelType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionCurrentVersion"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/functions/" + tea.StringValue(functionName) + "/current-version"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionCurrentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionDefaultInstance(appGroupIdentity *string, functionName *string) (_result *GetFunctionDefaultInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionDefaultInstanceResponse{}
	_body, _err := client.GetFunctionDefaultInstanceWithOptions(appGroupIdentity, functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionDefaultInstanceWithOptions(appGroupIdentity *string, functionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionDefaultInstanceResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionDefaultInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/default-instance"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionDefaultInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionInstance(appGroupIdentity *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest) (_result *GetFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.GetFunctionInstanceWithOptions(appGroupIdentity, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	instanceName = openapiutil.GetEncodeParam(instanceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Output)) {
		query["output"] = request.Output
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances/" + tea.StringValue(instanceName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionVersion(functionName *string, versionId *string) (_result *GetFunctionVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionVersionResponse{}
	_body, _err := client.GetFunctionVersionWithOptions(functionName, versionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionVersionWithOptions(functionName *string, versionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionVersionResponse, _err error) {
	functionName = openapiutil.GetEncodeParam(functionName)
	versionId = openapiutil.GetEncodeParam(versionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionVersion"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/functions/" + tea.StringValue(functionName) + "/versions/" + tea.StringValue(versionId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelProgress(appGroupIdentity *string, modelName *string) (_result *GetModelProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelProgressResponse{}
	_body, _err := client.GetModelProgressWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelProgressWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelProgressResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelProgress"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName) + "/progress"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelReport(appGroupIdentity *string, modelName *string) (_result *GetModelReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelReportResponse{}
	_body, _err := client.GetModelReportWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelReportWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelReportResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelReport"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName) + "/report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScriptFileNames(appGroupIdentity *string, appVersionId *string, scriptName *string) (_result *GetScriptFileNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScriptFileNamesResponse{}
	_body, _err := client.GetScriptFileNamesWithOptions(appGroupIdentity, appVersionId, scriptName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScriptFileNamesWithOptions(appGroupIdentity *string, appVersionId *string, scriptName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetScriptFileNamesResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetScriptFileNames"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName) + "/file-names"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScriptFileNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSearchStrategy(appGroupIdentity *string, appId *string, strategyName *string) (_result *GetSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSearchStrategyResponse{}
	_body, _err := client.GetSearchStrategyWithOptions(appGroupIdentity, appId, strategyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSearchStrategyWithOptions(appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSearchStrategyResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	strategyName = openapiutil.GetEncodeParam(strategyName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/search-strategies/" + tea.StringValue(strategyName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *GetSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSortScriptResponse{}
	_body, _err := client.GetSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSortScriptResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSortScriptFile(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string) (_result *GetSortScriptFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSortScriptFileResponse{}
	_body, _err := client.GetSortScriptFileWithOptions(appGroupIdentity, scriptName, appVersionId, fileName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSortScriptFileWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSortScriptFileResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	fileName = openapiutil.GetEncodeParam(fileName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSortScriptFile"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName) + "/files/src/" + tea.StringValue(fileName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSortScriptFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetValidationError(appGroupIdentity *string, request *GetValidationErrorRequest) (_result *GetValidationErrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetValidationErrorResponse{}
	_body, _err := client.GetValidationErrorWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetValidationErrorWithOptions(appGroupIdentity *string, request *GetValidationErrorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetValidationErrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		query["errorCode"] = request.ErrorCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetValidationError"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/data/validation-error"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetValidationErrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetValidationReport(appGroupIdentity *string, request *GetValidationReportRequest) (_result *GetValidationReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetValidationReportResponse{}
	_body, _err := client.GetValidationReportWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetValidationReportWithOptions(appGroupIdentity *string, request *GetValidationReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetValidationReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetValidationReport"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/data/validation-report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetValidationReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestExperiments(appGroupIdentity *string, sceneId *string, groupId *string) (_result *ListABTestExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestExperimentsResponse{}
	_body, _err := client.ListABTestExperimentsWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestExperimentsWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestExperimentsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestExperiments"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestFixedFlowDividers(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *ListABTestFixedFlowDividersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestFixedFlowDividersResponse{}
	_body, _err := client.ListABTestFixedFlowDividersWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestFixedFlowDividersWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestFixedFlowDividersResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	experimentId = openapiutil.GetEncodeParam(experimentId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestFixedFlowDividers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments/" + tea.StringValue(experimentId) + "/fixed-flow-dividers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestFixedFlowDividersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestGroups(appGroupIdentity *string, sceneId *string) (_result *ListABTestGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestGroupsResponse{}
	_body, _err := client.ListABTestGroupsWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestGroupsWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestGroupsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestGroups"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestMetrics(appGroupIdentity *string, sceneId *string, groupId *string) (_result *ListABTestMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestMetricsResponse{}
	_body, _err := client.ListABTestMetricsWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestMetricsWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestMetricsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestMetrics"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestScenes(appGroupIdentity *string) (_result *ListABTestScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestScenesResponse{}
	_body, _err := client.ListABTestScenesWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestScenesWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestScenesResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestScenes"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppGroupErrors(appGroupIdentity *string, request *ListAppGroupErrorsRequest) (_result *ListAppGroupErrorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppGroupErrorsResponse{}
	_body, _err := client.ListAppGroupErrorsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppGroupErrorsWithOptions(appGroupIdentity *string, request *ListAppGroupErrorsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppGroupErrorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StopTime)) {
		query["stopTime"] = request.StopTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppGroupErrors"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/errors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppGroupErrorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppGroupMetrics(appGroupIdentity *string, request *ListAppGroupMetricsRequest) (_result *ListAppGroupMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppGroupMetricsResponse{}
	_body, _err := client.ListAppGroupMetricsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppGroupMetricsWithOptions(appGroupIdentity *string, request *ListAppGroupMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppGroupMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Indexes)) {
		query["indexes"] = request.Indexes
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["metricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppGroupMetrics"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppGroupMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppGroups(request *ListAppGroupsRequest) (_result *ListAppGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppGroupsResponse{}
	_body, _err := client.ListAppGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppGroupsWithOptions(request *ListAppGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["sortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppGroups"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/apps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataCollections(appGroupIdentity *string, request *ListDataCollectionsRequest) (_result *ListDataCollectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataCollectionsResponse{}
	_body, _err := client.ListDataCollectionsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataCollectionsWithOptions(appGroupIdentity *string, request *ListDataCollectionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataCollectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataCollections"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/data-collections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSourceTableFields(dataSourceType *string, request *ListDataSourceTableFieldsRequest) (_result *ListDataSourceTableFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceTableFieldsResponse{}
	_body, _err := client.ListDataSourceTableFieldsWithOptions(dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceTableFieldsWithOptions(dataSourceType *string, request *ListDataSourceTableFieldsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceTableFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	dataSourceType = openapiutil.GetEncodeParam(dataSourceType)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTableFields"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/data-sources/" + tea.StringValue(dataSourceType) + "/fields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTableFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSourceTables(dataSourceType *string, request *ListDataSourceTablesRequest) (_result *ListDataSourceTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceTablesResponse{}
	_body, _err := client.ListDataSourceTablesWithOptions(dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceTablesWithOptions(dataSourceType *string, request *ListDataSourceTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	dataSourceType = openapiutil.GetEncodeParam(dataSourceType)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTables"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/data-sources/" + tea.StringValue(dataSourceType) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployedAlgorithmModels(appGroupIdentity *string, request *ListDeployedAlgorithmModelsRequest) (_result *ListDeployedAlgorithmModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDeployedAlgorithmModelsResponse{}
	_body, _err := client.ListDeployedAlgorithmModelsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployedAlgorithmModelsWithOptions(appGroupIdentity *string, request *ListDeployedAlgorithmModelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDeployedAlgorithmModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmType)) {
		query["algorithmType"] = request.AlgorithmType
	}

	if !tea.BoolValue(util.IsUnset(request.InServiceOnly)) {
		query["inServiceOnly"] = request.InServiceOnly
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployedAlgorithmModels"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/deployed-algorithm-models"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeployedAlgorithmModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFirstRanks(appGroupIdentity *string, appId *string) (_result *ListFirstRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFirstRanksResponse{}
	_body, _err := client.ListFirstRanksWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFirstRanksWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFirstRanksResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListFirstRanks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/first-ranks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFirstRanksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionInstances(appGroupIdentity *string, functionName *string, request *ListFunctionInstancesRequest) (_result *ListFunctionInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.ListFunctionInstancesWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionInstancesWithOptions(appGroupIdentity *string, functionName *string, request *ListFunctionInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		query["functionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["modelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		query["output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionInstances"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionTasks(appGroupIdentity *string, functionName *string, instanceName *string, request *ListFunctionTasksRequest) (_result *ListFunctionTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionTasksResponse{}
	_body, _err := client.ListFunctionTasksWithOptions(appGroupIdentity, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionTasksWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, request *ListFunctionTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	instanceName = openapiutil.GetEncodeParam(instanceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionTasks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances/" + tea.StringValue(instanceName) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaries(request *ListInterventionDictionariesRequest) (_result *ListInterventionDictionariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionariesResponse{}
	_body, _err := client.ListInterventionDictionariesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionariesWithOptions(request *ListInterventionDictionariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaryEntries(name *string, request *ListInterventionDictionaryEntriesRequest) (_result *ListInterventionDictionaryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionaryEntriesResponse{}
	_body, _err := client.ListInterventionDictionaryEntriesWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionaryEntriesWithOptions(name *string, request *ListInterventionDictionaryEntriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionaryEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		query["word"] = request.Word
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaryEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(name) + "/entries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionaryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaryNerResults(name *string, request *ListInterventionDictionaryNerResultsRequest) (_result *ListInterventionDictionaryNerResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionaryNerResultsResponse{}
	_body, _err := client.ListInterventionDictionaryNerResultsWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionaryNerResultsWithOptions(name *string, request *ListInterventionDictionaryNerResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionaryNerResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaryNerResults"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(name) + "/ner-results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionaryNerResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaryRelatedEntities(name *string) (_result *ListInterventionDictionaryRelatedEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionaryRelatedEntitiesResponse{}
	_body, _err := client.ListInterventionDictionaryRelatedEntitiesWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionaryRelatedEntitiesWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionaryRelatedEntitiesResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaryRelatedEntities"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(name) + "/related"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionaryRelatedEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModels(appGroupIdentity *string, request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsWithOptions(appGroupIdentity *string, request *ListModelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModels"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProceedings(appGroupIdentity *string) (_result *ListProceedingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProceedingsResponse{}
	_body, _err := client.ListProceedingsWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProceedingsWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProceedingsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListProceedings"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/proceedings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProceedingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueryProcessorAnalyzerResults(appGroupIdentity *string, appId *string, name *string, request *ListQueryProcessorAnalyzerResultsRequest) (_result *ListQueryProcessorAnalyzerResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryProcessorAnalyzerResultsResponse{}
	_body, _err := client.ListQueryProcessorAnalyzerResultsWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueryProcessorAnalyzerResultsWithOptions(appGroupIdentity *string, appId *string, name *string, request *ListQueryProcessorAnalyzerResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryProcessorAnalyzerResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryProcessorAnalyzerResults"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/query-processors/" + tea.StringValue(name) + "/analyze"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryProcessorAnalyzerResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueryProcessorNers(request *ListQueryProcessorNersRequest) (_result *ListQueryProcessorNersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryProcessorNersResponse{}
	_body, _err := client.ListQueryProcessorNersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueryProcessorNersWithOptions(request *ListQueryProcessorNersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryProcessorNersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryProcessorNers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/query-processor/ner/default-priorities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryProcessorNersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueryProcessors(appGroupIdentity *string, appId *string, request *ListQueryProcessorsRequest) (_result *ListQueryProcessorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryProcessorsResponse{}
	_body, _err := client.ListQueryProcessorsWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueryProcessorsWithOptions(appGroupIdentity *string, appId *string, request *ListQueryProcessorsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryProcessorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsActive)) {
		query["isActive"] = request.IsActive
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryProcessors"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/query-processors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryProcessorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuotaReviewTasks(appGroupIdentity *string, request *ListQuotaReviewTasksRequest) (_result *ListQuotaReviewTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotaReviewTasksResponse{}
	_body, _err := client.ListQuotaReviewTasksWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuotaReviewTasksWithOptions(appGroupIdentity *string, request *ListQuotaReviewTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotaReviewTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaReviewTasks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/quota-review-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaReviewTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRamRoles() (_result *ListRamRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRamRolesResponse{}
	_body, _err := client.ListRamRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRamRolesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRamRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRamRoles"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/ram/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRamRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScheduledTasks(appGroupIdentity *string, request *ListScheduledTasksRequest) (_result *ListScheduledTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.ListScheduledTasksWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScheduledTasksWithOptions(appGroupIdentity *string, request *ListScheduledTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListScheduledTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScheduledTasks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scheduled-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSearchStrategies(appGroupIdentity *string, appId *string) (_result *ListSearchStrategiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSearchStrategiesResponse{}
	_body, _err := client.ListSearchStrategiesWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchStrategiesWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSearchStrategiesResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSearchStrategies"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/search-strategies"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSearchStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecondRanks(appGroupIdentity *string, appId *string) (_result *ListSecondRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSecondRanksResponse{}
	_body, _err := client.ListSecondRanksWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecondRanksWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSecondRanksResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecondRanks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/second-ranks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecondRanksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlowQueryCategories(appGroupIdentity *string) (_result *ListSlowQueryCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSlowQueryCategoriesResponse{}
	_body, _err := client.ListSlowQueryCategoriesWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSlowQueryCategoriesWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSlowQueryCategoriesResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSlowQueryCategories"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/optimizers/slow-query/categories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSlowQueryCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlowQueryQueries(appGroupIdentity *string, categoryIndex *string) (_result *ListSlowQueryQueriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSlowQueryQueriesResponse{}
	_body, _err := client.ListSlowQueryQueriesWithOptions(appGroupIdentity, categoryIndex, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSlowQueryQueriesWithOptions(appGroupIdentity *string, categoryIndex *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSlowQueryQueriesResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	categoryIndex = openapiutil.GetEncodeParam(categoryIndex)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSlowQueryQueries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/optimizers/slow-query/categories/" + tea.StringValue(categoryIndex) + "/queries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSlowQueryQueriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSortExpressions(appGroupIdentity *string, appId *string) (_result *ListSortExpressionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSortExpressionsResponse{}
	_body, _err := client.ListSortExpressionsWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSortExpressionsWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSortExpressionsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSortExpressions"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/sort-expressions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSortExpressionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSortScripts(appGroupIdentity *string, appVersionId *string) (_result *ListSortScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSortScriptsResponse{}
	_body, _err := client.ListSortScriptsWithOptions(appGroupIdentity, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSortScriptsWithOptions(appGroupIdentity *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSortScriptsResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSortScripts"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSortScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStatisticLogs(appGroupIdentity *string, moduleName *string, request *ListStatisticLogsRequest) (_result *ListStatisticLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStatisticLogsResponse{}
	_body, _err := client.ListStatisticLogsWithOptions(appGroupIdentity, moduleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStatisticLogsWithOptions(appGroupIdentity *string, moduleName *string, request *ListStatisticLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStatisticLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	moduleName = openapiutil.GetEncodeParam(moduleName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Columns)) {
		query["columns"] = request.Columns
	}

	if !tea.BoolValue(util.IsUnset(request.Distinct)) {
		query["distinct"] = request.Distinct
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["sortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StopTime)) {
		query["stopTime"] = request.StopTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStatisticLogs"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/statistic-logs/" + tea.StringValue(moduleName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStatisticLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStatisticReport(appGroupIdentity *string, moduleName *string, request *ListStatisticReportRequest) (_result *ListStatisticReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStatisticReportResponse{}
	_body, _err := client.ListStatisticReportWithOptions(appGroupIdentity, moduleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStatisticReportWithOptions(appGroupIdentity *string, moduleName *string, request *ListStatisticReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStatisticReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	moduleName = openapiutil.GetEncodeParam(moduleName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Columns)) {
		query["columns"] = request.Columns
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStatisticReport"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/statistic-report/" + tea.StringValue(moduleName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStatisticReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserAnalyzerEntries(name *string, request *ListUserAnalyzerEntriesRequest) (_result *ListUserAnalyzerEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserAnalyzerEntriesResponse{}
	_body, _err := client.ListUserAnalyzerEntriesWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserAnalyzerEntriesWithOptions(name *string, request *ListUserAnalyzerEntriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserAnalyzerEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		query["word"] = request.Word
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserAnalyzerEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(name) + "/entries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserAnalyzerEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserAnalyzers(request *ListUserAnalyzersRequest) (_result *ListUserAnalyzersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserAnalyzersResponse{}
	_body, _err := client.ListUserAnalyzersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserAnalyzersWithOptions(request *ListUserAnalyzersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserAnalyzersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserAnalyzers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserAnalyzersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppGroup(appGroupIdentity *string) (_result *ModifyAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAppGroupResponse{}
	_body, _err := client.ModifyAppGroupWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppGroupWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAppGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppGroupQuota(appGroupIdentity *string) (_result *ModifyAppGroupQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAppGroupQuotaResponse{}
	_body, _err := client.ModifyAppGroupQuotaWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppGroupQuotaWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAppGroupQuotaResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppGroupQuota"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/quota"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppGroupQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFirstRank(appGroupIdentity *string, appId *string, name *string, request *ModifyFirstRankRequest) (_result *ModifyFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyFirstRankResponse{}
	_body, _err := client.ModifyFirstRankWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFirstRankWithOptions(appGroupIdentity *string, appId *string, name *string, request *ModifyFirstRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyFirstRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/first-ranks/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyModel(appGroupIdentity *string, modelName *string) (_result *ModifyModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyModelResponse{}
	_body, _err := client.ModifyModelWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyModelWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyModelResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyQueryProcessor(appGroupIdentity *string, appId *string, name *string, request *ModifyQueryProcessorRequest) (_result *ModifyQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyQueryProcessorResponse{}
	_body, _err := client.ModifyQueryProcessorWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyQueryProcessorWithOptions(appGroupIdentity *string, appId *string, name *string, request *ModifyQueryProcessorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyQueryProcessorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/query-processors/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScheduledTask(appGroupIdentity *string, taskId *string) (_result *ModifyScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.ModifyScheduledTaskWithOptions(appGroupIdentity, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScheduledTaskWithOptions(appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	taskId = openapiutil.GetEncodeParam(taskId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scheduled-tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecondRank(appGroupIdentity *string, appId *string, name *string, request *ModifySecondRankRequest) (_result *ModifySecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifySecondRankResponse{}
	_body, _err := client.ModifySecondRankWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecondRankWithOptions(appGroupIdentity *string, appId *string, name *string, request *ModifySecondRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifySecondRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/second-ranks/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreviewModel(appGroupIdentity *string, modelName *string, request *PreviewModelRequest) (_result *PreviewModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewModelResponse{}
	_body, _err := client.PreviewModelWithOptions(appGroupIdentity, modelName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreviewModelWithOptions(appGroupIdentity *string, modelName *string, request *PreviewModelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PreviewModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviewModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName) + "/actions/preview"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PreviewModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushInterventionDictionaryEntries(name *string) (_result *PushInterventionDictionaryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInterventionDictionaryEntriesResponse{}
	_body, _err := client.PushInterventionDictionaryEntriesWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushInterventionDictionaryEntriesWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInterventionDictionaryEntriesResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PushInterventionDictionaryEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(name) + "/entries/actions/bulk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushInterventionDictionaryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushUserAnalyzerEntries(name *string) (_result *PushUserAnalyzerEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushUserAnalyzerEntriesResponse{}
	_body, _err := client.PushUserAnalyzerEntriesWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushUserAnalyzerEntriesWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushUserAnalyzerEntriesResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PushUserAnalyzerEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(name) + "/entries/actions/bulk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushUserAnalyzerEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RankPreviewQuery(appGroupIdentity *string, modelName *string) (_result *RankPreviewQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RankPreviewQueryResponse{}
	_body, _err := client.RankPreviewQueryWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RankPreviewQueryWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RankPreviewQueryResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RankPreviewQuery"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName) + "/actions/query-rank"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RankPreviewQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *ReleaseSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseSortScriptResponse{}
	_body, _err := client.ReleaseSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseSortScriptResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName) + "/actions/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveApp(appGroupIdentity *string, appId *string) (_result *RemoveAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveAppResponse{}
	_body, _err := client.RemoveAppWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAppWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveAppResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveApp"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAppGroup(appGroupIdentity *string) (_result *RemoveAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveAppGroupResponse{}
	_body, _err := client.RemoveAppGroupWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAppGroupWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveAppGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDataCollection(appGroupIdentity *string, dataCollectionIdentity *string) (_result *RemoveDataCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveDataCollectionResponse{}
	_body, _err := client.RemoveDataCollectionWithOptions(appGroupIdentity, dataCollectionIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDataCollectionWithOptions(appGroupIdentity *string, dataCollectionIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveDataCollectionResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	dataCollectionIdentity = openapiutil.GetEncodeParam(dataCollectionIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDataCollection"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/data-collections/" + tea.StringValue(dataCollectionIdentity)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDataCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveFirstRank(appGroupIdentity *string, appId *string, name *string) (_result *RemoveFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveFirstRankResponse{}
	_body, _err := client.RemoveFirstRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveFirstRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveFirstRankResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/first-ranks/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveInterventionDictionary(name *string) (_result *RemoveInterventionDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveInterventionDictionaryResponse{}
	_body, _err := client.RemoveInterventionDictionaryWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveInterventionDictionaryWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveInterventionDictionaryResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveInterventionDictionary"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveInterventionDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveQueryProcessor(appGroupIdentity *string, appId *string, name *string) (_result *RemoveQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveQueryProcessorResponse{}
	_body, _err := client.RemoveQueryProcessorWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveQueryProcessorWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveQueryProcessorResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/query-processors/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveScheduledTask(appGroupIdentity *string, taskId *string) (_result *RemoveScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveScheduledTaskResponse{}
	_body, _err := client.RemoveScheduledTaskWithOptions(appGroupIdentity, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveScheduledTaskWithOptions(appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveScheduledTaskResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	taskId = openapiutil.GetEncodeParam(taskId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scheduled-tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSearchStrategy(appGroupIdentity *string, appId *string, strategyName *string) (_result *RemoveSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveSearchStrategyResponse{}
	_body, _err := client.RemoveSearchStrategyWithOptions(appGroupIdentity, appId, strategyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSearchStrategyWithOptions(appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveSearchStrategyResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	strategyName = openapiutil.GetEncodeParam(strategyName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/search-strategies/" + tea.StringValue(strategyName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSecondRank(appGroupIdentity *string, appId *string, name *string) (_result *RemoveSecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveSecondRankResponse{}
	_body, _err := client.RemoveSecondRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSecondRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveSecondRankResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/second-ranks/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserAnalyzer(name *string) (_result *RemoveUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveUserAnalyzerResponse{}
	_body, _err := client.RemoveUserAnalyzerWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserAnalyzerWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveUserAnalyzerResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewAppGroup(appGroupIdentity *string) (_result *RenewAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewAppGroupResponse{}
	_body, _err := client.RenewAppGroupWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewAppGroupWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewAppGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RenewAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/actions/renew"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceAppGroupCommodityCode(appGroupIdentity *string) (_result *ReplaceAppGroupCommodityCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReplaceAppGroupCommodityCodeResponse{}
	_body, _err := client.ReplaceAppGroupCommodityCodeWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceAppGroupCommodityCodeWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReplaceAppGroupCommodityCodeResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceAppGroupCommodityCode"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/actions/to-instance-typed"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceAppGroupCommodityCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSortScriptFile(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string) (_result *SaveSortScriptFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveSortScriptFileResponse{}
	_body, _err := client.SaveSortScriptFileWithOptions(appGroupIdentity, scriptName, appVersionId, fileName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSortScriptFileWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveSortScriptFileResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	fileName = openapiutil.GetEncodeParam(fileName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSortScriptFile"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName) + "/files/src/" + tea.StringValue(fileName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSortScriptFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSlowQueryAnalyzer(appGroupIdentity *string) (_result *StartSlowQueryAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartSlowQueryAnalyzerResponse{}
	_body, _err := client.StartSlowQueryAnalyzerWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartSlowQueryAnalyzerWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartSlowQueryAnalyzerResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartSlowQueryAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/optimizers/slow-query/actions/run"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSlowQueryAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrainModel(appGroupIdentity *string, modelName *string) (_result *TrainModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TrainModelResponse{}
	_body, _err := client.TrainModelWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrainModelWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TrainModelResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	modelName = openapiutil.GetEncodeParam(modelName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("TrainModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/algorithm/models/" + tea.StringValue(modelName) + "/actions/train"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrainModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindESUserAnalyzer(appGroupIdentity *string, esInstanceId *string) (_result *UnbindESUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindESUserAnalyzerResponse{}
	_body, _err := client.UnbindESUserAnalyzerWithOptions(appGroupIdentity, esInstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindESUserAnalyzerWithOptions(appGroupIdentity *string, esInstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindESUserAnalyzerResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	esInstanceId = openapiutil.GetEncodeParam(esInstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindESUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/es/" + tea.StringValue(esInstanceId) + "/actions/unbind-analyzer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindESUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindEsInstance(appGroupIdentity *string) (_result *UnbindEsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindEsInstanceResponse{}
	_body, _err := client.UnbindEsInstanceWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindEsInstanceWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindEsInstanceResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEsInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/actions/unbind-es-instance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *UpdateABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestExperimentResponse{}
	_body, _err := client.UpdateABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestExperimentResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	experimentId = openapiutil.GetEncodeParam(experimentId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments/" + tea.StringValue(experimentId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestFixedFlowDividers(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *UpdateABTestFixedFlowDividersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestFixedFlowDividersResponse{}
	_body, _err := client.UpdateABTestFixedFlowDividersWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestFixedFlowDividersWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestFixedFlowDividersResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	experimentId = openapiutil.GetEncodeParam(experimentId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestFixedFlowDividers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId) + "/experiments/" + tea.StringValue(experimentId) + "/fixed-flow-dividers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestFixedFlowDividersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestGroup(appGroupIdentity *string, sceneId *string, groupId *string) (_result *UpdateABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestGroupResponse{}
	_body, _err := client.UpdateABTestGroupWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestGroupResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	groupId = openapiutil.GetEncodeParam(groupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId) + "/groups/" + tea.StringValue(groupId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestScene(appGroupIdentity *string, sceneId *string) (_result *UpdateABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestSceneResponse{}
	_body, _err := client.UpdateABTestSceneWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestSceneWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestSceneResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	sceneId = openapiutil.GetEncodeParam(sceneId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/scenes/" + tea.StringValue(sceneId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFetchFields(appGroupIdentity *string, appId *string, request *UpdateFetchFieldsRequest) (_result *UpdateFetchFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFetchFieldsResponse{}
	_body, _err := client.UpdateFetchFieldsWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFetchFieldsWithOptions(appGroupIdentity *string, appId *string, request *UpdateFetchFieldsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFetchFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFetchFields"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/fetch-fields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFetchFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunctionDefaultInstance(appGroupIdentity *string, functionName *string, request *UpdateFunctionDefaultInstanceRequest) (_result *UpdateFunctionDefaultInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionDefaultInstanceResponse{}
	_body, _err := client.UpdateFunctionDefaultInstanceWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionDefaultInstanceWithOptions(appGroupIdentity *string, functionName *string, request *UpdateFunctionDefaultInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionDefaultInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunctionDefaultInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/default-instance"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionDefaultInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunctionInstance(appGroupIdentity *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest) (_result *UpdateFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.UpdateFunctionInstanceWithOptions(appGroupIdentity, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	functionName = openapiutil.GetEncodeParam(functionName)
	instanceName = openapiutil.GetEncodeParam(instanceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateParameters)) {
		body["createParameters"] = request.CreateParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Cron)) {
		body["cron"] = request.Cron
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/functions/" + tea.StringValue(functionName) + "/instances/" + tea.StringValue(instanceName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSearchStrategy(appGroupIdentity *string, appId *string, strategyName *string) (_result *UpdateSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSearchStrategyResponse{}
	_body, _err := client.UpdateSearchStrategyWithOptions(appGroupIdentity, appId, strategyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSearchStrategyWithOptions(appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSearchStrategyResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	strategyName = openapiutil.GetEncodeParam(strategyName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/search-strategies/" + tea.StringValue(strategyName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSortScript(appGroupIdentity *string, appVersionId *string, scriptName *string) (_result *UpdateSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSortScriptResponse{}
	_body, _err := client.UpdateSortScriptWithOptions(appGroupIdentity, appVersionId, scriptName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSortScriptWithOptions(appGroupIdentity *string, appVersionId *string, scriptName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSortScriptResponse, _err error) {
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appVersionId = openapiutil.GetEncodeParam(appVersionId)
	scriptName = openapiutil.GetEncodeParam(scriptName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appVersionId) + "/sort-scripts/" + tea.StringValue(scriptName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSummaries(appGroupIdentity *string, appId *string, request *UpdateSummariesRequest) (_result *UpdateSummariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSummariesResponse{}
	_body, _err := client.UpdateSummariesWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSummariesWithOptions(appGroupIdentity *string, appId *string, request *UpdateSummariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSummariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appGroupIdentity = openapiutil.GetEncodeParam(appGroupIdentity)
	appId = openapiutil.GetEncodeParam(appId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSummaries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(appGroupIdentity) + "/apps/" + tea.StringValue(appId) + "/summaries"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSummariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateDataSources() (_result *ValidateDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateDataSourcesResponse{}
	_body, _err := client.ValidateDataSourcesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateDataSourcesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateDataSourcesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateDataSources"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/data-sources/validations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}