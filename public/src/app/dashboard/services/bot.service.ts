import { catchError, map } from 'rxjs/operators';
import { Bot, BuildStatus, Volume, Env, Config } from './../models/bot';
import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
} from '@angular/common/http';
import { throwError, Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { RunStatus } from '../models/bot';

export const BAD_FORMAT = 'bad format';
export const NOT_FOUND = 'not found';
export const CONFLICT = 'conflict';

@Injectable({
  providedIn: 'root',
})
export class BotService {
  constructor(private http: HttpClient) {}

  options = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
    }),
  };

  private handleError(error: HttpErrorResponse): Observable<never> {
    if (error.status === 404) {
      return throwError(NOT_FOUND);
    } else if (error.status === 409) {
      return throwError(CONFLICT);
    } else if (error.status === 400) {
      return throwError(BAD_FORMAT);
    }

    console.error(error);
    return throwError(error.message);
  }

  getVolumes(id: number): Observable<Volume[]> {
    return this.http
      .get<Volume[]>(`${environment.apiUrl}/bot/${id}/volume`)
      .pipe(catchError(this.handleError));
  }

  addVolume(id: number, vol: Volume): Observable<void> {
    return this.http
      .post(`${environment.apiUrl}/bot/${id}/volume`, vol, this.options)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  deleteVolume(id: number, vol: string): Observable<void> {
    return this.http
      .delete(`${environment.apiUrl}/bot/${id}/volume/${vol}`)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  getConfigs(id: number): Observable<Config[]> {
    return this.http
      .get<Config[]>(`${environment.apiUrl}/bot/${id}/config`)
      .pipe(catchError(this.handleError));
  }

  addConfig(id: number, conf: Config): Observable<void> {
    return this.http
      .post(`${environment.apiUrl}/bot/${id}/config`, conf, this.options)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  deleteConfig(id: number, conf: string): Observable<void> {
    return this.http
      .delete(`${environment.apiUrl}/bot/${id}/config/${conf}`)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  getEnvs(id: number): Observable<Env[]> {
    return this.http
      .get<Env[]>(`${environment.apiUrl}/bot/${id}/env`)
      .pipe(catchError(this.handleError));
  }

  addEnv(id: number, env: Env): Observable<void> {
    return this.http
      .post(`${environment.apiUrl}/bot/${id}/env`, env, this.options)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  deleteEnv(id: number, env: string): Observable<void> {
    return this.http.delete(`${environment.apiUrl}/bot/${id}/env/${env}`).pipe(
      catchError(this.handleError),
      map(_ => {})
    );
  }

  getBots(): Observable<Bot[]> {
    return this.http
      .get<Bot[]>(`${environment.apiUrl}/bot`)
      .pipe(catchError(this.handleError));
  }

  rebuildBot(id: number): Observable<void> {
    return this.http
      .post(`${environment.apiUrl}/bot/${id}/build`, this.options)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  uploadBot(id: number): Observable<void> {
    return this.http
      .put(`${environment.apiUrl}/bot/${id}/run`, this.options)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  getBuildStatus(id: number): Observable<BuildStatus> {
    return this.http
      .get<BuildStatus>(`${environment.apiUrl}/bot/${id}/status/build`)
      .pipe(catchError(this.handleError));
  }

  getRunStatus(id: number): Observable<RunStatus> {
    return this.http
      .get<RunStatus>(`${environment.apiUrl}/bot/${id}/status/run`)
      .pipe(catchError(this.handleError));
  }
}
