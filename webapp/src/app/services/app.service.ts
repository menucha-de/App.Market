import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable } from 'rxjs';
import { App } from '../models/app';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  private baseUrl = 'rest/';
  installing: Map<string, boolean>;
  private configUrl: URL;

  constructor(private http: HttpClient) {
    this.configUrl = new URL(window.location.href);
    this.installing = new Map<string, boolean>();
  }

  getInstalledApps(namespace: string) {
    return this.http.get<App[]>(`${this.baseUrl}apps/${namespace}/installed`);
  }

  getAllApps(namespace: string) {
    return this.http.get<App[]>(`${this.baseUrl}apps/${namespace}`);
  }

  installApp(namespace: string, app: App): Observable<void> {
    return this.http.post<void>(`${this.baseUrl}apps/${namespace}`, app);
  }

  upgradeApps(namespace: string, apps: App[]): Observable<void> {
    return this.http.put<void>(`${this.baseUrl}apps/${namespace}`, apps);
  }

  updateApp(namespace: string, name: string, app: App): Observable<void> {
    return this.http.put<void>(`${this.baseUrl}apps/${namespace}/${name}`, app);
  }

  deleteApp(namespace: string, name: string): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}apps/${namespace}/${name}`);
  }
  uploadImgFile(namespace: string, file: File, name: string) {
    const headers = new HttpHeaders().append('Content-Type', 'application/octet-stream');
    return this.http.put<void>(`${this.baseUrl}apps/${namespace}/add/${name}`, file, { headers });
  }

  sort(apps: App[]): App[] {
    apps.sort((a: App, b: App) => {
      if (a.label.toLowerCase() > b.label.toLowerCase()) {
        return 1;
      }
      if (a.label.toLowerCase() < b.label.toLowerCase()) {
        return -1;
      }
      return 0;
    });

    return apps;
  }
}
