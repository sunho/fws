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
    }

    console.error(error);
    return throwError(error.message);
  }

  getVolumes(id: number): Observable<Volume[]> {
    return this.http
      .get<Volume[]>(`${environment.apiUrl}/bot/${id}/volume`)
      .pipe(catchError(this.handleError));
  }

  getConfigs(id: number): Observable<Config[]> {
    return this.http
      .get<Config[]>(`${environment.apiUrl}/bot/${id}/config`)
      .pipe(catchError(this.handleError));
  }

  getEnvs(id: number): Observable<Env[]> {
    return this.http
      .get<Env[]>(`${environment.apiUrl}/bot/${id}/env`)
      .pipe(catchError(this.handleError));
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

  getBotBuildStatus(id: number): Observable<BuildStatus> {
    return this.http
      .get<BuildStatus>(`${environment.apiUrl}/bot/${id}/status/build`)
      .pipe(catchError(this.handleError));
  }
}
