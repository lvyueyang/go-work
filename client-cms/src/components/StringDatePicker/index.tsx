import { DatePicker } from 'antd';
import { PickerPanelDateProps } from 'antd/es/calendar/generateCalendar';
import { DatePickerType } from 'antd/es/date-picker';
import dayjs from 'dayjs';
import { ComponentProps, forwardRef } from 'react';

export const StringDatePicker = forwardRef<
  any,
  Omit<ComponentProps<typeof DatePicker> & PickerPanelDateProps<DatePickerType>, 'onChange'> & {
    onChange?: (value: string | null, dateString: string | null) => void;
  }
>((props, ref) => {
  const val = typeof props.value === 'string' ? dayjs(props.value) : props.value;
  return (
    <DatePicker
      {...props}
      value={val}
      ref={ref}
      showTime
      onChange={(e) => {
        props.onChange?.(e?.format('YYYY-MM-DD HH:mm:ss') ?? null, e?.toString() ?? null);
      }}
    ></DatePicker>
  );
});
