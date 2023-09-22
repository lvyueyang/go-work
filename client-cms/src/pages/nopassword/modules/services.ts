import { AIP_FIX } from '@/constants';
import { ApiAdminUserForgetPasswordBodyDto } from '@/interface/serverApi';
import { Result, request } from '@/request';

export const forgetPassword = (body: ApiAdminUserForgetPasswordBodyDto) => {
  return request.post<Result<void>>(`${AIP_FIX}/auth/forget-password`, body);
};
