import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Urls } from '../../shared/constants/urls';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { getDefaultHttpOptions } from './default-http-options';
import { BaseString } from '../../types/base-string';

@Injectable({
    providedIn: 'root',
})
export class BaseStringService {
    url = `${environment.urlApi}/${Urls.baseStrings}`;

    constructor(private httpClient: HttpClient) {}

    create(baseString: BaseString): Observable<BaseString> {
        return this.httpClient.post<BaseString>(this.url, baseString, getDefaultHttpOptions());
    }

    findAll(): Observable<BaseString[]> {
        return this.httpClient
            .get<BaseString[]>(this.url, getDefaultHttpOptions())
            .pipe(map((baseStrings) => (baseStrings ? baseStrings : [])));
    }

    update(baseString: BaseString): Observable<BaseString> {
        return this.httpClient.put<BaseString>(this.url, baseString, getDefaultHttpOptions());
    }
}
