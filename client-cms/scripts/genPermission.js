const http = require('http');
const fs = require('fs');
const path = require('path');

http
  .get(`http://127.0.0.1:4002/api/admin/role/permission/codes`, (response) => {
    let dataStr = '';

    response.on('data', (chunk) => {
      dataStr += chunk;
    });

    response.on('end', () => {
      const data = JSON.parse(dataStr);
      console.log(data.data);
      let codeResult = '';
      let mapResult = '';
      Object.entries(data.data).forEach(([k, item]) => {
        const key = k.toLocaleUpperCase().replace(/:/gi, '_');
        const label = item.label;
        codeResult += `  /** ${label} */
  ${key}: '${k}',
`;
        mapResult += `  [PERMISSION_CODE.${key}]: '${label}',
`;
      });

      template(codeResult, mapResult);
    });
  })
  .on('error', (error) => {
    console.log('Error: ' + error.message);
  });

function template(codeResult, mapResult) {
  const result = `export const PERMISSION_CODE = {
${codeResult}} as const;
  
export const PERMISSION_MAP = {
${mapResult}} as const;
`;

  console.log(result);
  fs.writeFileSync(path.join(__dirname, '../src/permission/codes.ts'), result);
}
