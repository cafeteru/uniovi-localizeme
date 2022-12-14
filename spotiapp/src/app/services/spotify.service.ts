import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import { map, switchMap, tap } from 'rxjs/operators';
import { environment } from 'src/environments/environment';

interface LoginResponse {
    access_token: string;
    token_type: string;
    expires_in: number;
}

@Injectable({
    providedIn: 'root',
})
export class SpotifyService {
    constructor(private httpClient: HttpClient) {}

    getToken(): Observable<LoginResponse> {
        const url = 'https://accounts.spotify.com/api/token?Authorization';
        const urlencoded = new URLSearchParams();
        urlencoded.append('grant_type', 'client_credentials');
        urlencoded.append('client_id', environment.clientId);
        urlencoded.append('client_secret', environment.clientSecret);
        return this.httpClient
            .post<LoginResponse>(url, urlencoded, {
                responseType: 'json',
                headers: new HttpHeaders({
                    'Content-Type': 'application/x-www-form-urlencoded',
                }),
            })
            .pipe(tap((loginResponse) => localStorage.setItem('authorization', loginResponse.access_token)));
    }

    getQuery(query: string): Observable<any> {
        const url = `https://api.spotify.com/v1/${query}`;
        return this.getToken().pipe(
            switchMap(() =>
                this.httpClient.get(url, {
                    responseType: 'json',
                    headers: new HttpHeaders({
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${localStorage.authorization}`,
                    }),
                })
            )
        );
    }

    getNewReleases(): Observable<any> {
        return this.getQuery('browse/new-releases?limit=20').pipe(map((data) => data['albums'].items));
    }

    getArtists(termino: string): Observable<any> {
        return this.getQuery(`search?q=${termino}&type=artist&limit=15`).pipe(map((data) => data['artists'].items));
    }

    getArtist(id: string): Observable<any> {
        return this.getQuery(`artists/${id}`);
    }

    getTopTracks(id: string): Observable<any> {
        return this.getQuery(`artists/${id}/top-tracks?country=us`).pipe(map((data) => data['tracks']));
    }
}
