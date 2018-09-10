import { catchError, map } from 'rxjs/operators';
import { AppConfig } from './../../app.config';
import { Bot, BuildStatus } from './../models/bot';
import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
} from '@angular/common/http';
import { throwError, Observable } from 'rxjs';

export const NOT_FOUND = 'not found';

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
    }

    console.error(error);
    return throwError('unknown error');
  }

  getBots(): Observable<Bot[]> {
    return this.http
      .get<Bot[]>(`${AppConfig.apiUrl}/bot`)
      .pipe(catchError(this.handleError));
  }

  rebuildBot(id: number): Observable<void> {
    return this.http
      .post(`${AppConfig.apiUrl}/bot/${id}/build`, this.options)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  getBotBuildStatus(id: number): Observable<BuildStatus> {
    return this.http
      .get<BuildStatus>(`${AppConfig.apiUrl}/bot/${id}/status/build`)
      .pipe(catchError(this.handleError));
  }
}
