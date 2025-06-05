import { AIP_FIX } from '@/constants';
import { ApiAdminInitRootUserReq, ApiAdminInitRootUserRes } from '@opd/api-interface';
import { Result, request } from '@/request';

export const initRootUser = (body: ApiAdminInitRootUserReq) => {
  return request.post<Result<ApiAdminInitRootUserRes>>(`${AIP_FIX}/auth/init-root-user`, body);
};
