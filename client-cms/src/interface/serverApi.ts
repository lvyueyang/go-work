/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ApiAdminInitRootUserReq {
  /** 邮箱 */
  email: string;
  /** 昵称 */
  name: string;
  /** 密码 */
  password: string;
  /** 用户名 */
  username: string;
}

export interface ApiAdminInitRootUserRes {
  token?: string;
}

export interface ApiAdminRoleCreateReq {
  /** 角色编码 */
  code: string;
  /** 描述 */
  desc?: string;
  /** 角色名称 */
  name: string;
}

export interface ApiAdminRoleCreateRes {
  /** 角色 ID */
  id?: number;
}

export interface ApiAdminRoleDeleteReq {
  /** 角色 ID */
  id: number;
}

export interface ApiAdminRoleListReq {
  /** @min 1 */
  current?: number;
  /** 查询关键字 */
  keyword?: string;
  order_key?: string;
  order_type?: 'ase' | 'desc';
  /** @min 0 */
  page_size?: number;
}

export interface ApiAdminRoleListRes {
  list?: ModelAdminRole[];
  total?: number;
}

export interface ApiAdminRoleUpdateCodesReq {
  /** 权限码 */
  codes: string[];
  /** 角色 ID */
  id: number;
}

export interface ApiAdminRoleUpdateReq {
  /** 角色编码 */
  code: string;
  /** 描述 */
  desc?: string;
  /** 角色 ID */
  id: number;
  /** 角色名称 */
  name: string;
}

export interface ApiAdminUserCreateReq {
  /** 头像 */
  avatar?: string;
  /** 邮箱 */
  email: string;
  /** 姓名 */
  name: string;
  /** 密码 */
  password: string;
  /** 用户名 */
  username: string;
}

export interface ApiAdminUserCreateRes {
  /** 管理员ID */
  id?: string;
}

export interface ApiAdminUserDeleteReq {
  /** ID */
  id: number;
}

export interface ApiAdminUserForgetPasswordReq {
  /** 邮箱验证码 */
  captcha: string;
  /** 邮箱 */
  email: string;
  /** 密码 */
  password: string;
}

export interface ApiAdminUserInfoRes {
  avatar?: string;
  /** 创建时间 */
  created_at: string;
  email?: string;
  /** ID */
  id: number;
  /** 是否是超级管理员 */
  is_root?: boolean;
  name: string;
  news?: ModelNews[];
  roles?: ModelAdminRole[];
  status: ConstsAdminUserStatus;
  /** 更新时间 */
  updated_at?: string;
  username: string;
}

export interface ApiAdminUserListReq {
  /** @min 1 */
  current?: number;
  /** 查询关键字 */
  keyword?: string;
  order_key?: string;
  order_type?: 'ase' | 'desc';
  /** @min 0 */
  page_size?: number;
}

export interface ApiAdminUserListRes {
  list?: ModelAdminUser[];
  total?: number;
}

export interface ApiAdminUserLoginReq {
  /** 密码 */
  password?: string;
  /** 用户名 */
  username?: string;
}

export interface ApiAdminUserLoginRes {
  token?: string;
}

export interface ApiAdminUserResetPasswordReq {
  /** ID */
  id: number;
  /** 密码 */
  password: string;
}

export interface ApiAdminUserUpdateReq {
  /** 头像 */
  avatar?: string;
  /** ID */
  id: number;
  /** 姓名 */
  name: string;
}

export interface ApiAdminUserUpdateRoleReq {
  /** 角色 ID */
  role_ids: number[];
  /** 用户 ID */
  user_id: number;
}

export interface ApiAdminUserUpdateStatusReq {
  /** ID */
  id: number;
  /** 状态 1-解封 2-封禁 */
  status: -1 | 1;
}

export interface ApiNewsCreateReq {
  content?: string;
  cover?: string;
  desc?: string;
  /** 可见性 */
  is_visible?: boolean;
  /** 发布日期 YYYY-MM-DD HH:mm:ss */
  push_date?: string;
  /** 推荐等级 0 为不推荐，数值越大越靠前 */
  recommend?: number;
  title: string;
}

export interface ApiNewsCreateRes {
  /** 新闻 ID */
  id: number;
}

export interface ApiNewsInfoRes {
  /** 作者 */
  author_id: number;
  /** 内容 */
  content?: string;
  /** 封面 */
  cover?: string;
  /** 创建时间 */
  created_at: string;
  /** 简介 */
  desc?: string;
  /** ID */
  id: number;
  /** 可见性 */
  is_visible: boolean;
  /** 发布日期 */
  push_date?: string;
  /** 推荐等级 */
  recommend: number;
  /** 新闻标题 */
  title: string;
  /** 更新时间 */
  updated_at?: string;
}

export interface ApiNewsListReq {
  /** @min 1 */
  current?: number;
  /** 查询关键字 */
  keyword?: string;
  order_key?: string;
  order_type?: 'ase' | 'desc';
  /** @min 0 */
  page_size?: number;
}

export interface ApiNewsListRes {
  list?: ModelNews[];
  total?: number;
}

export interface ApiNewsUpdateReq {
  content?: string;
  cover?: string;
  desc?: string;
  /** 新闻 ID */
  id: number;
  /** 可见性 */
  is_visible?: boolean;
  /** 发布日期 YYYY-MM-DD HH:mm:ss */
  push_date?: string;
  /** 推荐等级 0 为不推荐，数值越大越靠前 */
  recommend?: number;
  title: string;
}

export enum ConstsAdminUserStatus {
  AdminUserStatusNormal = 1,
  AdminUserStatusLocked = -1,
}

export enum ConstsCaptchaScenes {
  CaptchaScenesRegister = 1,
  CaptchaScenesForgetPassword = 2,
  CaptchaScenesUpdatePhone = 3,
  CaptchaScenesUpdateEmail = 4,
}

export enum ConstsCaptchaType {
  CaptchaTypePhone = 1,
  CaptchaTypeEmail = 2,
}

export enum ConstsUserStatus {
  UserStatusNormal = 1,
  UserStatusLocked = -1,
}

export interface ControllerCreateCaptchaBodyDto {
  /** 图形验证码的key */
  captcha_key: string;
  /** 输入的图形验证码 */
  captcha_value: string;
  /** 使用场景， 1-注册 2-忘记密码 3-修改手机 4-修改邮箱 */
  scenes: ConstsCaptchaScenes;
  /** 验证码类型， 1-手机 2-邮箱 */
  type: ConstsCaptchaType;
  /** 手机/邮箱账号 */
  value: string;
}

export interface ControllerUpdateUserStatusBodyDto {
  /** 用户 ID */
  id: number;
  /** 状态 -1封禁 1-正常 */
  status: -1 | 1;
}

export interface ControllerLoginBodyDto {
  /** 密码 */
  password?: string;
  /** 用户名 */
  username?: string;
}

export interface ControllerLoginSuccessResponse {
  token?: string;
}

export interface ControllerRegisterBodyDto {
  /** 邮箱验证码 */
  captcha?: string;
  /** 邮箱 */
  email: string;
  /** 密码 */
  password: string;
  /** 用户名 */
  username: string;
}

export interface ControllerWxMpLoginBodyDto {
  /** code */
  code: string;
}

export interface ModelAdminRole {
  /** 角色编号 */
  code: string;
  /** 创建时间 */
  created_at: string;
  desc?: string;
  /** ID */
  id: number;
  name: string;
  /** 权限码 */
  permission_codes?: string[];
  /** 更新时间 */
  updated_at?: string;
  users: ModelAdminUser[];
}

export interface ModelAdminUser {
  avatar?: string;
  /** 创建时间 */
  created_at: string;
  email?: string;
  /** ID */
  id: number;
  /** 是否是超级管理员 */
  is_root?: boolean;
  name: string;
  news?: ModelNews[];
  roles?: ModelAdminRole[];
  status: ConstsAdminUserStatus;
  /** 更新时间 */
  updated_at?: string;
  username: string;
}

export interface ModelNews {
  /** 作者 */
  author_id: number;
  /** 内容 */
  content?: string;
  /** 封面 */
  cover?: string;
  /** 创建时间 */
  created_at: string;
  /** 简介 */
  desc?: string;
  /** ID */
  id: number;
  /** 可见性 */
  is_visible: boolean;
  /** 发布日期 */
  push_date?: string;
  /** 推荐等级 */
  recommend: number;
  /** 新闻标题 */
  title: string;
  /** 更新时间 */
  updated_at?: string;
}

export interface ModelUser {
  age?: number;
  avatar?: string;
  /** 创建时间 */
  created_at: string;
  email: string;
  /** ID */
  id: number;
  name?: string;
  status: ConstsUserStatus;
  /** 更新时间 */
  updated_at?: string;
  username: string;
  wx_open_id?: string;
  wx_union_id?: string;
}

export interface RespRList {
  list?: any[];
  total?: number;
}

export interface RespResult {
  code?: number;
  data?: any;
  msg?: string;
}

export interface UtilsPermissionInfo {
  /** 权限码 */
  code: string;
  /** 权限名称 */
  name: string;
}
