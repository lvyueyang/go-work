interface Size {
  width: number;
  height: number;
}

/** 获取图片长度和宽度 */
export function getImageSize(file: File | string): Promise<Size> {
  const img = new Image();
  if (typeof file === 'string') {
    img.src = file;
  } else {
    img.src = URL.createObjectURL(file);
  }
  return new Promise((resolve, reject) => {
    img.onload = () => {
      console.log(img.width, img.height);
      resolve({
        width: img.width,
        height: img.height,
      });
    };
    img.onerror = () => {
      reject('图片加载失败');
    };
  });
}

export interface ValidateRule {
  minWidth?: number;
  minHeight?: number;
  maxWidth?: number;
  maxHeight?: number;
  width?: number;
  height?: number;
  /** 宽高比 16/9 => [16,9] */
  scale?: [number, number];
}

/** 求最小公约数 */
function getGcd(a: number, b: number): number {
  if (b === 0) {
    return a;
  }
  return getGcd(b, a % b);
}

/** 校验尺寸 */
export function validateSize(size: Size, rule: ValidateRule) {
  console.log(size);
  if (typeof rule.width === 'number' && size.width !== rule.width) {
    return `宽度不正确，当前为 ${size.width} px，应为 ${rule.width} px`;
  }
  if (typeof rule.height === 'number' && size.height !== rule.height) {
    return `高度不正确，当前为 ${size.height} px，应为 ${rule.height} px`;
  }
  if (typeof rule.minWidth === 'number' && size.width < rule.minWidth) {
    return `宽度不正确，当前为 ${size.width} px，最小应为 ${rule.minWidth} px`;
  }
  if (typeof rule.minHeight === 'number' && size.height < rule.minHeight) {
    return `高度不正确，当前为 ${size.height} px，最小应为 ${rule.minHeight} px`;
  }
  if (typeof rule.maxWidth === 'number' && size.width > rule.maxWidth) {
    return `宽度不正确，当前为 ${size.width} px，最大应为 ${rule.maxWidth} px`;
  }
  if (typeof rule.maxHeight === 'number' && size.height > rule.maxHeight) {
    return `高度不正确，当前为 ${size.height} px，最大应为 ${rule.maxHeight} px`;
  }
  if (Array.isArray(rule.scale)) {
    const gcd = getGcd(rule.scale[0], rule.scale[1]);
    if (size.width / gcd !== rule.scale[0] || size.height / gcd !== rule.scale[1]) {
      return `比例不正确，当前为 ${size.width / gcd} / ${size.height / gcd}，应为 ${rule.scale.join(
        '/',
      )}`;
    }
  }
  return false;
}
