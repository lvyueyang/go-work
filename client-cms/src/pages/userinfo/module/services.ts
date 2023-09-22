import { AIP_FIX, Result, request } from '@/request';
import { UpdatePasswordBody } from './types';

/** 修改密码 */
export const updatePassword = (body: UpdatePasswordBody) => {
  return request.post<Result<void>>(`${AIP_FIX}/user/UpdateUserInfo`, body);
};
