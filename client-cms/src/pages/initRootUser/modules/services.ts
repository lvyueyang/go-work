import { AIP_FIX } from '@/constants';
import {
  ApiAdminInitRootUserBodyDto,
  ApiAdminUserLoginSuccessResponse,
} from '@/interface/serverApi';
import { Result, request } from '@/request';

export const initRootUser = (body: ApiAdminInitRootUserBodyDto) => {
  return request.post<Result<ApiAdminUserLoginSuccessResponse>>(
    `${AIP_FIX}/auth/init-root-user`,
    body,
  );
};
