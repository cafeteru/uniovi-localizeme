import { Observable, of } from 'rxjs';
import { LoginData, User } from '../../../types/user';

export class UserServiceMock {
    delete(user: User): Observable<boolean> {
        return of(Boolean(user));
    }

    disable(user: User): Observable<User> {
        return of(user);
    }

    findAll(): Observable<User[]> {
        return of(undefined);
    }

    findMe(): Observable<User> {
        return of(undefined);
    }

    login(loginData: LoginData): Observable<User> {
        return of({
            id: '',
            admin: false,
            active: true,
            email: loginData.email,
            password: loginData.password,
        });
    }

    update(user: User): Observable<User> {
        return of(user);
    }

    updateMe(user: User): Observable<User> {
        return of(user);
    }

    register(loginData: LoginData): Observable<User> {
        return of({
            id: '',
            admin: false,
            active: true,
            email: loginData.email,
            password: loginData.password,
        });
    }
}
