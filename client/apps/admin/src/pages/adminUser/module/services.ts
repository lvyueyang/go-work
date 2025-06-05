import { AIP_FIX } from '@/constants';
import {
  ApiAdminUserCreateReq,
  ApiAdminUserCreateRes,
  ApiAdminUserListReq,
  ApiAdminUserListRes,
  ApiAdminUserResetPasswordReq,
  ApiAdminUserUpdateReq,
  ApiAdminUserUpdateRoleReq,
  ApiAdminUserUpdateStatusReq,
  ModelAdminUser,
} from '@opd/api-interface';
import { Result, request } from '@/request';

/** 用户列表 */
export const getUserListApi = (body: ApiAdminUserListReq) => {
  return request.post<Result<ApiAdminUserListRes>>(`${AIP_FIX}/user/list`, body);
};

/** 详情 */
export const getUserDetail = (id: number) => {
  return request.post<Result<ModelAdminUser>>(`${AIP_FIX}/user/${id}`);
};

/** 创建 */
export const createUserApi = (body: ApiAdminUserCreateReq) => {
  return request.post<Result<ApiAdminUserCreateRes>>(`${AIP_FIX}/user/create`, body);
};

/** 修改 */
export const updateUserApi = (body: ApiAdminUserUpdateReq) => {
  return request.put<Result<void>>(`${AIP_FIX}/user/update/info`, body);
};

/** 修改角色 */
export const updateRoleApi = (body: ApiAdminUserUpdateRoleReq) => {
  return request.post<Result<void>>(`${AIP_FIX}/user/update/role`, body);
};

/** 重置密码 */
export const resetPasswordApi = (body: ApiAdminUserResetPasswordReq) => {
  return request.post<Result<void>>(`${AIP_FIX}/user/reset-password`, body);
};

/** 封禁/解封 */
export const updateStatusApi = (body: ApiAdminUserUpdateStatusReq) => {
  return request.post<Result<void>>(`${AIP_FIX}/user/update/status`, body);
};

/** 删除用户 */
export const deleteUserApi = (id: number) => {
  return request.post<Result<void>>(`${AIP_FIX}/user/delete`, { id });
};
