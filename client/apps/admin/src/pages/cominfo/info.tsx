import Editor from '@/components/Editor';
import { COMINFO_TYPE } from '@/constants';
import { CominfoUpdateDto } from '@opd/api-interface';
import { Button, Card, Form, Row } from 'antd';
import { useComInfo } from './module/useInfo';

type FormValues = CominfoUpdateDto;

export default function CominfoIntroduceForm() {
  const { form, submitLoading, submitHandler } = useComInfo(COMINFO_TYPE.JING_ZHNEG_LI.value);

  return (
    <>
      <Card style={{ maxWidth: 1000, margin: '0 auto' }}>
        <Form<FormValues> form={form} onFinish={submitHandler} labelCol={{ xs: 3 }}>
          <Form.Item name="content" rules={[{ required: true, validateTrigger: 'submit' }]}>
            <Editor placeholder="请填写公司核心竞争力" style={{ height: 700 }} />
          </Form.Item>
          <Form.Item>
            <Row justify="center">
              <Button style={{ width: 160 }} type="primary" htmlType="submit" loading={submitLoading}>
                提交
              </Button>
            </Row>
          </Form.Item>
        </Form>
      </Card>
    </>
  );
}
