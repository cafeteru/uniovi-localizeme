import { createMockUser, User } from './user';
import { Permission } from './permission';

export interface GroupDto {
    name: string;
    owner: User;
    permissions: Permission[];
    public: boolean;
}

export interface Group {
    id: string;
    name: string;
    owner: User;
    permissions: Permission[];
    active: boolean;
    public: boolean;
}

export function createMockGroup(): Group {
    return {
        id: '1',
        name: 'group',
        owner: createMockUser(),
        active: true,
        permissions: [],
        public: true,
    };
}
