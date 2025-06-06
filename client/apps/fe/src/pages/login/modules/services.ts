import { ApiAdminUserLoginBodyDto, ApiAdminUserLoginSuccessResponse } from '@opd/api-interface';
import { AIP_FIX, Result, request } from '@/request';

export const login = (body: ApiAdminUserLoginBodyDto) => {
  return request.post<Result<ApiAdminUserLoginSuccessResponse>>(`${AIP_FIX}/auth/login`, body);
};
