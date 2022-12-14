import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserInfoComponent } from './user-info.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { CoreModule } from '../../../core/core.module';
import { SharedModule } from '../../../shared/shared.module';
import { provideMockStore } from '@ngrx/store/testing';
import { AppState } from '../../../store/app.reducer';
import { createMockAppState } from '../../../store/mocks/create-mock-app-state';

describe('UserInfoComponent', () => {
    let component: UserInfoComponent;
    let fixture: ComponentFixture<UserInfoComponent>;
    let appState: AppState;

    beforeEach(async () => {
        appState = createMockAppState();
        await TestBed.configureTestingModule({
            declarations: [UserInfoComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule],
            providers: [provideMockStore({ initialState: appState })],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(UserInfoComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
