import { AIP_FIX } from '@/constants';
import { ListResult, Result, request } from '@/request';
import {
  ApiCreateNewsBodyDto,
  ApiUpdateNewsBodyDto,
  ModelNews,
  ServiceFindNewsListOption,
} from 'interface/serverApi';

/** 列表 */
export const getListApi = (body: ServiceFindNewsListOption) => {
  return request.post<ListResult<ModelNews>>(`${AIP_FIX}/news/list`, body);
};

/** 详情 */
export const getDetailApi = (id: number) => {
  return request.get<Result<ModelNews>>(`${AIP_FIX}/news/${id}`);
};

/** 创建 */
export const createApi = (body: ApiCreateNewsBodyDto) => {
  return request.post<Result<void>>(`${AIP_FIX}/news`, body);
};

/** 修改 */
export const updateApi = (body: ApiUpdateNewsBodyDto) => {
  return request.put<Result<void>>(`${AIP_FIX}/news/`, body);
};

/** 删除 */
export const removeApi = (id: number | string) => {
  return request.delete<Result<void>>(`${AIP_FIX}/news/${id}`);
};
