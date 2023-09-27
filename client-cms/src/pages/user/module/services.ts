import { ADMIN_USER_STATUS_ENUM, AIP_FIX } from '@/constants';
import { ListResult, Pagination, Result, request } from '@/request';
import { ModelUser } from 'interface/serverApi';

/** 用户列表 */
export const getUserList = (query: Pagination & { keyword?: string }) => {
  return request.get<ListResult<ModelUser>>(`${AIP_FIX}/c-user`, {
    params: {
      ...query,
    },
  });
};

/** 封禁/解封 */
export const updateStatus = (id: number, status: ADMIN_USER_STATUS_ENUM) => {
  return request.put<Result<void>>(`${AIP_FIX}/c-user/status`, {
    status,
    id,
  });
};
