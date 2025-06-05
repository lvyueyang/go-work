import { AIP_FIX } from '@/constants';
import {
  ApiNewsCreateReq,
  ApiNewsCreateRes,
  ApiNewsInfoRes,
  ApiNewsListReq,
  ApiNewsListRes,
  ApiNewsUpdateReq,
} from '@opd/api-interface';
import { Result, request } from '@/request';
/** 列表 */
export const getListApi = (body: ApiNewsListReq) => {
  return request.post<Result<ApiNewsListRes>>(`${AIP_FIX}/news/list`, body);
};

/** 详情 */
export const getDetailApi = (id: number) => {
  return request.post<Result<ApiNewsInfoRes>>(`${AIP_FIX}/news/info`, { id });
};

/** 创建 */
export const createApi = (body: ApiNewsCreateReq) => {
  return request.post<Result<ApiNewsCreateRes>>(`${AIP_FIX}/news/create`, body);
};

/** 修改 */
export const updateApi = (body: ApiNewsUpdateReq) => {
  return request.post<Result<void>>(`${AIP_FIX}/news/update/info`, body);
};

/** 删除 */
export const removeApi = (id: number) => {
  return request.post<Result<void>>(`${AIP_FIX}/news/delete`, { id });
};
