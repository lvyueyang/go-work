import { COMINFO_TYPE_ENUM } from '@/constants';
import { CominfoUpdateDto } from '@/interface/serverApi';
import { message } from '@/utils/notice';
import { useRequest } from 'ahooks';
import { Form } from 'antd';
import { useEffect, useState } from 'react';
import { getDetailApi, updateApi } from '.';

type FormValues = CominfoUpdateDto;

export function useComInfo(type: COMINFO_TYPE_ENUM) {
  const [submitLoading, setSubmitLoading] = useState(false);
  const [form] = Form.useForm<FormValues>();
  const { data: info } = useRequest(() => {
    return getDetailApi(type).then((res) => res.data.data);
  });
  const submitHandler = (values: FormValues) => {
    setSubmitLoading(true);
    return updateApi(type, values)
      .then(() => {
        message.success('修改成功');
      })
      .finally(() => {
        setSubmitLoading(false);
      });
  };
  useEffect(() => {
    if (info) {
      form.setFieldsValue(info);
    }
  }, [info]);
  return { info, form, submitLoading, submitHandler };
}
