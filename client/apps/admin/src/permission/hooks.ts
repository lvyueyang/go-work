import { useUserinfoStore } from '@/store/userinfo';
import { TypeValue } from '@/types';
import { PERMISSION_CODE } from '.';

export type CodeType = TypeValue<typeof PERMISSION_CODE>;

export function usePermission(...codes: CodeType[]) {
  const { data, permissions } = useUserinfoStore();
  if (data?.is_root) {
    return true;
  }
  // 差集
  const difference = new Set([...codes].filter((x) => !permissions.has(x)));
  return difference.size === 0;
}
