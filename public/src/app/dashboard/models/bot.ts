export interface Bot {
  id: number;
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
