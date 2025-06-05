import { ProTable, ProTableProps } from '@ant-design/pro-components';
import type { ParamsType } from '@ant-design/pro-provider';
function PageTable<
  DataType extends Record<string, any>,
  Params extends ParamsType = ParamsType,
  ValueType = 'text',
>(props: ProTableProps<DataType, Params, ValueType>) {
  return (
    <ProTable<DataType, Params, ValueType>
      search={false}
      bordered
      scroll={{ y: window.innerHeight - 305 }}
      rowKey="id"
      {...props}
    />
  );
}

export default PageTable;
