import { getImageCaptchaCode } from '@/services';
import { useRequest } from 'ahooks';
import { Input, Space } from 'antd';
import { InputProps } from 'antd/lib/input';

export interface ImageCaptchaValue {
  value: string;
  key: string;
}

interface ImageCaptchaProps {
  inputProps?: InputProps;
  value?: ImageCaptchaValue;
  onChange?: (value: ImageCaptchaValue) => void;
}

export function ImageCaptcha({ value, onChange, inputProps }: ImageCaptchaProps) {
  const { data, run } = useRequest(() => {
    return getImageCaptchaCode().then((res) => {
      const key = res.data.data;
      onChange?.({
        key,
        value: '',
      });
      return key;
    });
  });
  return (
    <Space>
      <Input
        value={value?.value}
        onChange={(e) => {
          onChange?.({
            key: data!,
            value: e.target.value,
          });
        }}
        {...inputProps}
      />
      <img
        src={`/api/captcha/image/${data}`}
        alt=""
        height={40}
        style={{ cursor: 'pointer' }}
        onClick={() => {
          run();
        }}
      />
    </Space>
  );
}
