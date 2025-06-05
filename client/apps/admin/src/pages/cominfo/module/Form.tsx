import Editor from '@/components/Editor';
import UploadImage from '@/components/UploadImage';
import { CominfoUpdateDto } from '@opd/api-interface';
import { COMINFO_TYPE_ENUM } from '@/constants';
import { Button, Card, Form, Input, Row } from 'antd';
import { useComInfo } from './useInfo';

type FormValues = CominfoUpdateDto;

export function ComInfoForm({ type }: { type: COMINFO_TYPE_ENUM }) {
  const { info, form, submitHandler } = useComInfo(type);

  return (
    <Card style={{ maxWidth: 1000, margin: '0 auto' }}>
      <Form<FormValues> form={form} onFinish={submitHandler} labelCol={{ xs: 3 }}>
        <Form.Item label="标题" name="title" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="封面" name="cover" rules={[{ required: true, message: '请上传封面' }]}>
          <UploadImage />
        </Form.Item>
        <Form.Item label="描述" name="desc">
          <Input.TextArea />
        </Form.Item>
        <Form.Item label="详情" name="content" rules={[{ required: true, validateTrigger: 'submit' }]}>
          <Editor style={{ height: 400 }} />
        </Form.Item>
        <Form.Item label=" ">
          <Row justify="center">
            <Button style={{ width: 160 }} type="primary" htmlType="submit" loading={loading}>
              提交
            </Button>
          </Row>
        </Form.Item>
      </Form>
    </Card>
  );
}
