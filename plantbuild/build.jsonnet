local c = import 'dc.jsonnet';
local dc = c {
  dockerRegistry: '562055475000.dkr.ecr.ap-northeast-1.amazonaws.com',
};

dc.build_apps_image('qor5/qor5', [
  { name: 'docs', dockerfile: './Dockerfile', context: '.' },
])