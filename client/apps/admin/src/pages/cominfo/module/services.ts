import { AIP_FIX, COMINFO_TYPE_ENUM } from '@/constants';
import { CominfoDetailResponseDto, CominfoUpdateDto } from '@opd/api-interface';
import { Result, request } from '@/request';

/** 详情 */
export const getDetailApi = (type: COMINFO_TYPE_ENUM) => {
  return request.get<CominfoDetailResponseDto>(`${AIP_FIX}/cominfo/${type}`);
};

/** 修改 */
export const updateApi = (type: COMINFO_TYPE_ENUM, body: CominfoUpdateDto) => {
  return request.put<Result<string>>(`${AIP_FIX}/cominfo/${type}`, body);
};
