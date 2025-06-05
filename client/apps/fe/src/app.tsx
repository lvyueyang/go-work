import 'antd/dist/reset.css';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import advancedFormat from 'dayjs/plugin/advancedFormat';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import localeData from 'dayjs/plugin/localeData';
import weekOfYear from 'dayjs/plugin/weekOfYear';
import weekYear from 'dayjs/plugin/weekYear';
import weekday from 'dayjs/plugin/weekday';
import React from 'react';
import { RuntimeConfig } from 'umi';
import pkg from '../package.json';
import { ThemeProvider } from './theme';

dayjs.extend(customParseFormat);
dayjs.extend(advancedFormat);
dayjs.extend(weekday);
dayjs.extend(localeData);
dayjs.extend(weekOfYear);
dayjs.extend(weekYear);

dayjs.locale('zh-cn');
console.log('version:', pkg.version);

export const rootContainer: RuntimeConfig['rootContainer'] = (container) => {
  return React.createElement(ThemeProvider, null, container);
};
