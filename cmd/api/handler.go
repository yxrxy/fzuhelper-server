/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	api "github.com/west2-online/fzuhelper-server/kitex_gen/api"
)

// CommonServiceImpl implements the last service interface defined in the IDL.
type CommonServiceImpl struct{}

// GetCSS implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetCSS(ctx context.Context, req *api.GetCSSRequest) (resp *api.GetCSSResponse, err error) {
	// TODO: Your code here...
	return
}

// GetHtml implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetHtml(ctx context.Context, req *api.GetHtmlRequest) (resp *api.GetHtmlResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserAgreement implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetUserAgreement(ctx context.Context, req *api.GetUserAgreementRequest) (resp *api.GetUserAgreementResponse, err error) {
	// TODO: Your code here...
	return
}

// GetTermsList implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetTermsList(ctx context.Context, req *api.TermListRequest) (resp *api.TermListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetTerm implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetTerm(ctx context.Context, req *api.TermRequest) (resp *api.TermResponse, err error) {
	// TODO: Your code here...
	return
}

// GetNotice implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetNotice(ctx context.Context, req *api.GetNoticeRequst) (resp *api.GetNoticeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetContributorInfo implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetContributorInfo(ctx context.Context, req *api.GetContributorInfoRequest) (resp *api.GetContributorInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetToolboxConfig implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) GetToolboxConfig(ctx context.Context, req *api.GetToolboxConfigRequest) (resp *api.GetToolboxConfigResponse, err error) {
	// TODO: Your code here...
	return
}

// PutToolboxConfig implements the CommonServiceImpl interface.
func (s *CommonServiceImpl) PutToolboxConfig(ctx context.Context, req *api.PutToolboxConfigRequest) (resp *api.PutToolboxConfigResponse, err error) {
	// TODO: Your code here...
	return
}
