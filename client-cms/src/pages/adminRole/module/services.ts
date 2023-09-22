import { AIP_FIX } from '@/constants';
import {
  ApiCreateAdminRoleBodyDto,
  ApiUpdateAdminRoleBodyDto,
  ApiUpdateAdminRolePermissionBodyDto,
  ModelAdminRole,
  PermissionLabelType,
} from '@/interface/serverApi';
import { ListResult, Pagination, Result, request } from '@/request';

/** 权限码列表 */
export const getCodeListApi = () => {
  return request.get<Result<Record<string, Required<PermissionLabelType>>>>(
    `${AIP_FIX}/role/permission/codes`,
  );
};

/** 修改权限 */
export const updateCodeApi = (body: ApiUpdateAdminRolePermissionBodyDto) => {
  return request.put<Result<void>>(`${AIP_FIX}/role/permission/codes`, body);
};

/** 列表 */
export const getListApi = (params: Pagination & { keyword?: string }) => {
  return request.get<ListResult<ModelAdminRole>>(`${AIP_FIX}/role`, {
    params,
  });
};

/** 创建 */
export const createApi = (body: ApiCreateAdminRoleBodyDto) => {
  return request.post<Result<void>>(`${AIP_FIX}/role`, body);
};

/** 修改 */
export const updateApi = (body: ApiUpdateAdminRoleBodyDto) => {
  return request.put<Result<void>>(`${AIP_FIX}/role`, body);
};

/** 删除 */
export const removeApi = (id: number | string) => {
  return request.delete<Result<void>>(`${AIP_FIX}/role/${id}`);
};
