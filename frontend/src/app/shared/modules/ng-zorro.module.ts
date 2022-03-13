import { NgModule } from '@angular/core';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageServiceModule } from 'ng-zorro-antd/message';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzAvatarModule } from 'ng-zorro-antd/avatar';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { NzPopoverModule } from 'ng-zorro-antd/popover';

@NgModule({
    imports: [
        NzAvatarModule,
        NzButtonModule,
        NzCheckboxModule,
        NzDropDownModule,
        NzFormModule,
        NzIconModule,
        NzInputModule,
        NzLayoutModule,
        NzMenuModule,
        NzMessageServiceModule,
        NzModalModule,
        NzPopoverModule,
        NzSelectModule,
    ],
    exports: [
        NzAvatarModule,
        NzButtonModule,
        NzCheckboxModule,
        NzDropDownModule,
        NzFormModule,
        NzIconModule,
        NzInputModule,
        NzLayoutModule,
        NzMenuModule,
        NzMessageServiceModule,
        NzModalModule,
        NzPopoverModule,
        NzSelectModule,
    ],
})
export class NgZorroModule {}
