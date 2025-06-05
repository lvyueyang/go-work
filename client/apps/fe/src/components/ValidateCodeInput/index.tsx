import { SEND_TYPE_ENUM, SEND_VALIDATE_CODE_TYPE_ENUM } from '@/constants';
import { sendCaptcha } from '@/services';
import { message } from '@/utils/notice';
import { Button, Input, InputProps, Popover, Space } from 'antd';
import { useRef, useState } from 'react';
import { ImageCaptcha, ImageCaptchaValue } from '../Captcha/ImageCaptcha';

const DEF_TEST = '发送验证码';

interface SendButtonProps {
  targetValue: string;
  sendType: SEND_VALIDATE_CODE_TYPE_ENUM;
  actionType: SEND_TYPE_ENUM;
}
export function SendButton({ targetValue, sendType, actionType }: SendButtonProps) {
  const [captcha, setCaptcha] = useState<ImageCaptchaValue>({
    key: '',
    value: '',
  });
  const [pop, setPop] = useState(false);
  const [loading, setLoading] = useState(false);
  const [disabled, setDisabled] = useState(false);
  const [text, setText] = useState(DEF_TEST);
  const timer = useRef(0);

  const looperTimer = () => {
    if (timer.current === 0) {
      setText(DEF_TEST);
      setDisabled(false);
      return;
    }
    setTimeout(() => {
      timer.current -= 1;
      const newText = `${timer.current}s 后重新发送`;
      setText(newText);

      looperTimer();
    }, 1000);
  };

  const submitHandler = async () => {
    if (loading) return;
    try {
      setLoading(true);
      setText('发送中 ...');
      await sendCaptcha({
        captcha_key: captcha.key,
        captcha_value: captcha.value,
        scenes: actionType,
        type: sendType,
        value: targetValue,
      });
      message.success('发送成功');
      setPop(false);
      timer.current = 60;
      setDisabled(true);
      looperTimer();
    } catch (e) {
      setText(DEF_TEST);
    }
    setLoading(false);
  };

  return (
    <Popover
      open={pop}
      title="请完成图形验证码验证"
      trigger={!targetValue || disabled ? [] : ['click']}
      content={
        <>
          <ImageCaptcha
            inputProps={{ style: { width: 120 } }}
            value={captcha}
            onChange={setCaptcha}
          />
          <div>
            <Space style={{ marginTop: 10 }}>
              <Button onClick={() => setPop(false)}>取消</Button>
              <Button type="primary" onClick={submitHandler}>
                确定
              </Button>
            </Space>
          </div>
        </>
      }
    >
      <Button
        type="primary"
        loading={loading}
        disabled={!targetValue || disabled}
        formNoValidate
        onClick={() => {
          setPop(true);
        }}
      >
        {text}
      </Button>
    </Popover>
  );
}

export interface ValidateCodeInputProps extends SendButtonProps, InputProps {}

export default function ValidateCodeInput({
  sendType,
  targetValue,
  placeholder = '请输入验证码',
  actionType,
  ...props
}: ValidateCodeInputProps) {
  return (
    <div style={{ display: 'flex', gap: 10 }}>
      <Input type="text" placeholder={placeholder} {...props} />
      <SendButton targetValue={targetValue} sendType={sendType} actionType={actionType} />
    </div>
  );
}
