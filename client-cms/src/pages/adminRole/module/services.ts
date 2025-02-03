import { AIP_FIX } from '@/constants';
import {
  ApiAdminRoleCreateReq,
  ApiAdminRoleCreateRes,
  ApiAdminRoleListReq,
  ApiAdminRoleListRes,
  ApiAdminRoleUpdateCodesReq,
  ApiAdminRoleUpdateReq,
  UtilsPermissionInfo,
} from '@/interface/serverApi';
import { Result, request } from '@/request';

/** 列表 */
export const getListApi = (params: ApiAdminRoleListReq) => {
  return request.post<Result<ApiAdminRoleListRes>>(`${AIP_FIX}/role/list`, {
    params,
  });
};

/** 创建 */
export const createApi = (body: ApiAdminRoleCreateReq) => {
  return request.post<Result<ApiAdminRoleCreateRes>>(`${AIP_FIX}/role/create`, body);
};

/** 修改 */
export const updateApi = (body: ApiAdminRoleUpdateReq) => {
  return request.post<Result<void>>(`${AIP_FIX}/role/update`, body);
};

/** 删除 */
export const removeApi = (id: number) => {
  return request.post<Result<void>>(`${AIP_FIX}/role/delete`, { id });
};

/** 权限码列表 */
export const getCodeListApi = () => {
  return request.get<Result<UtilsPermissionInfo[]>>(`${AIP_FIX}/role/permission/codes`);
};

/** 修改权限 */
export const updateCodeApi = (body: ApiAdminRoleUpdateCodesReq) => {
  return request.post<Result<void>>(`${AIP_FIX}/role/update/permission-codes`, body);
};
