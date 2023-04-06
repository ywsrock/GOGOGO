import { faker } from '@faker-js/faker';
import { sample } from 'lodash';

// ----------------------------------------------------------------------

faker.setLocale('ja')
const users = [...Array(24)].map((_, index) => ({
  id: faker.datatype.uuid(),
  avatarUrl: `/assets/images/avatars/avatar_${index + 1}.jpg`,
  name: faker.name.fullName(),
  company: faker.company.name(),
  isVerified: sample(['2022-01-01', '2023-11-11']),
  status: sample(['active', 'banned']),
  role: sample([
    '注文問い合わせ',
    '配達時間...',
    '電話番号変更',
    '支払カード変更',
    '手数料',
  ]),
}));

export default users;
