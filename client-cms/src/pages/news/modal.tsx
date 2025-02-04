import Editor from '@/components/Editor';
import { StringDatePicker } from '@/components/StringDatePicker';
import UploadImage from '@/components/UploadImage';
import { NEWS_TYPE } from '@/constants';
import { useFormModal } from '@/hooks/useFormModal';
import { Button, Flex, Form, Input, InputNumber, Modal } from 'antd';

export function FormModal<T>({ formModal }: { formModal: ReturnType<typeof useFormModal<T>> }) {
  return (
    <Modal
      open={formModal.formModal.open}
      onCancel={formModal.formModalClose}
      title={formModal.formModalTitle + '新闻'}
      width={800}
      style={{ top: 10 }}
    >
      <br />
      <Form<T>
        form={formModal.form}
        onFinish={formModal.submitHandler}
        labelCol={{ xs: 3 }}
        initialValues={{ recommend: 0, type: NEWS_TYPE.COM.value }}
      >
        <Form.Item label="新闻标题" name="title" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="发布日期" name="push_date">
          <StringDatePicker format={'YYYY-MM-DD HH:mm'} showTime />
        </Form.Item>
        <Form.Item
          label="推荐等级"
          name="recommend"
          help="0 为不推荐，大于 0 会在首页根据值进行排序展示，值越大排列越靠前"
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item label="新闻封面" name="cover">
          <UploadImage />
        </Form.Item>
        <Form.Item label="新闻描述" name="desc">
          <Input.TextArea />
        </Form.Item>
        <Form.Item
          label="新闻详情"
          name="content"
          rules={[{ required: true, validateTrigger: 'submit' }]}
        >
          <Editor style={{ height: 400 }} />
        </Form.Item>
        <Form.Item label=" ">
          <Flex justify="center">
            <Button
              style={{ width: 160 }}
              type="primary"
              htmlType="submit"
              loading={formModal.submitLoading}
            >
              提交
            </Button>
          </Flex>
        </Form.Item>
      </Form>
    </Modal>
  );
}
