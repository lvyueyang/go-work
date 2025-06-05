import { CodeType, usePermission } from './hooks';

interface PermissionProps {
  code: CodeType | CodeType[];
}

export function Permission({ code, children }: React.PropsWithChildren<PermissionProps>) {
  const codes = Array.isArray(code) ? code : [code];
  const success = usePermission(...codes);
  if (success) {
    return <>{children}</>;
  }
  return null;
}
