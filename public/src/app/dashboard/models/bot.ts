export interface Bot {
  id: number;
  version: number;
  name: string;
  build_result: string;
  webhook_secret: string;
  git_url: string;
}

export interface BuildStatusBuilt {
  type: 'built';
  created: string;
}

export interface BuildStatusBuilding {
  type: 'built';
}

export interface BuildStatusNone {
  type: 'none';
}

export type BuildStatus =
  | BuildStatusBuilding
  | BuildStatusBuilt
  | BuildStatusNone;

export interface RunStatusNone {
  type: 'none';
}


export interface RunStatusRunning {
  type: 'built';
  seconds: number;
}

export interface RunStatusFailed {
  type: 'failed';
  seconds: number;
}

export interface RunStatusPending {
  type: 'pending';
}

export type RunStatus =
  RunStatusNone |
  RunStatusRunning |
  RunStatusFailed |
  RunStatusPending;

export interface Volume {
  bot_id?: number;
  size?: number;
  name: string;
  version: number;
  path: string;
}

export interface Config {
  bot_id?: number;
  value?: string;
  name: string;
  version: number;
  path: string;
  file: string;
}

export interface Env {
  bot_id?: number;
  value?: string;
  version: number;
  name: string;
}
