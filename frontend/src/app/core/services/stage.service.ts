import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { catchError, map, Observable, of } from 'rxjs';
import { Stage, StageDto } from '../../types/stage';
import { getDefaultHttpOptions } from './default-http-options';
import { Urls } from '../../shared/constants/urls';

@Injectable({
    providedIn: 'root',
})
export class StageService {
    url = `${environment.urlApi}/${Urls.stages}`;

    constructor(private httpClient: HttpClient) {}

    create(stageRequest: StageDto): Observable<Stage> {
        return this.httpClient.post<Stage>(this.url, stageRequest, getDefaultHttpOptions());
    }

    delete(stage: Stage): Observable<boolean> {
        return this.httpClient.delete<Stage>(`${this.url}/${stage.id}`, getDefaultHttpOptions()).pipe(
            map(() => true),
            catchError(() => of(false))
        );
    }

    disable(stage: Stage): Observable<Stage> {
        return this.httpClient.patch<Stage>(`${this.url}/${stage.id}`, stage, getDefaultHttpOptions());
    }

    findAll(): Observable<Stage[]> {
        return this.httpClient
            .get<Stage[]>(this.url, getDefaultHttpOptions())
            .pipe(map((stages) => (stages ? stages : [])));
    }

    update(stage: Stage): Observable<Stage> {
        return this.httpClient.put<Stage>(this.url, stage, getDefaultHttpOptions());
    }
}