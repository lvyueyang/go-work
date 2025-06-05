import { ADMIN_USER_STATUS_ENUM, AIP_FIX } from '@/constants';
import { Result, request } from '@/request';
import { ApiUserListReq, ApiUserListRes } from '@opd/api-interface';

/** 用户列表 */
export const getUserListApi = (body: ApiUserListReq) => {
  return request.post<Result<ApiUserListRes>>(`${AIP_FIX}/c-user/list`, body);
};

/** 封禁/解封 */
export const updateStatusApi = (id: number, status: ADMIN_USER_STATUS_ENUM) => {
  return request.post<Result<void>>(`${AIP_FIX}/c-user/update/status`, {
    status,
    id,
  });
};
