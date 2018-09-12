import { environment } from './../../../environments/environment';
import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpErrorResponse,
  HttpHeaders,
} from '@angular/common/http';

import { catchError, map } from 'rxjs/operators';
import { throwError, Observable } from 'rxjs';

import { User } from '../models/user';

export const NOT_FOUND = 'not found';
export const DUPLICATE = 'duplicate';
export const WRONG_CRED = 'wrong cred';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
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

    if (error.status === 403) {
      return throwError(WRONG_CRED);
    }

    if (error.status === 409) {
      return throwError(DUPLICATE);
    }

    console.error(error);
    return throwError(error.message);
  }

  getUser(): Observable<User> {
    return this.http
      .get<User>(`${environment.apiUrl}/user`, this.options)
      .pipe(catchError(this.handleError));
  }

  login(username: string, password: string): Observable<void> {
    return this.http
      .post(
        `${environment.apiUrl}/login`,
        { username: username, password: password },
        this.options
      )
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  logout(): Observable<void> {
    return this.http.delete(`${environment.apiUrl}/login`).pipe(
      catchError(this.handleError),
      map(_ => {})
    );
  }

  register(
    key: string,
    username: string,
    nickname: string,
    password: string
  ): Observable<void> {
    return this.http
      .post(
        `${environment.apiUrl}/register`,
        {
          key: key,
          username: username,
          nickname: nickname,
          password: password,
        },
        this.options
      )
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }

  keyCheck(key: string, username: string): Observable<void> {
    return this.http
      .get(`${environment.apiUrl}/invite/${username}?key=${key}`)
      .pipe(
        catchError(this.handleError),
        map(_ => {})
      );
  }
}
