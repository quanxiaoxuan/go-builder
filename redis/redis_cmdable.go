package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// redis命令窗
type Cmdable struct {
	config *Config
	cmd    redis.Cmdable
}

func (r *Cmdable) GetConfig() *Config {
	return r.config
}

func (r *Cmdable) Pipeline() redis.Pipeliner {
	return r.cmd.Pipeline()
}

func (r *Cmdable) Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return r.cmd.Pipelined(fn)
}

func (r *Cmdable) TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return r.cmd.TxPipelined(fn)
}

func (r *Cmdable) TxPipeline() redis.Pipeliner {
	return r.cmd.TxPipeline()
}

func (r *Cmdable) Command() *redis.CommandsInfoCmd {
	return r.cmd.Command()
}

func (r *Cmdable) ClientGetName() *redis.StringCmd {
	return r.cmd.ClientGetName()
}

func (r *Cmdable) Echo(message interface{}) *redis.StringCmd {
	return r.cmd.Echo(message)
}

func (r *Cmdable) Ping() *redis.StatusCmd {
	return r.cmd.Ping()
}
func (r *Cmdable) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.cmd.HScan(key, cursor, match, count)
}

func (r *Cmdable) Quit() *redis.StatusCmd {
	return r.cmd.Quit()
}

func (r *Cmdable) Del(keys ...string) *redis.IntCmd {
	return r.cmd.Del(keys...)
}

func (r *Cmdable) Unlink(keys ...string) *redis.IntCmd {
	return r.cmd.Unlink(keys...)
}

func (r *Cmdable) Dump(key string) *redis.StringCmd {
	return r.cmd.Dump(key)
}

func (r *Cmdable) Exists(keys ...string) *redis.IntCmd {
	return r.cmd.Exists(keys...)
}

func (r *Cmdable) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return r.cmd.Expire(key, expiration)
}

func (r *Cmdable) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return r.cmd.ExpireAt(key, tm)
}

func (r *Cmdable) Keys(pattern string) *redis.StringSliceCmd {
	return r.cmd.Keys(pattern)
}

func (r *Cmdable) Migrate(host, port, key string, db int64, timeout time.Duration) *redis.StatusCmd {
	return r.cmd.Migrate(host, port, key, db, timeout)
}

func (r *Cmdable) Move(key string, db int64) *redis.BoolCmd {
	return r.cmd.Move(key, db)
}

func (r *Cmdable) ObjectRefCount(key string) *redis.IntCmd {
	return r.cmd.ObjectRefCount(key)
}

func (r *Cmdable) ObjectEncoding(key string) *redis.StringCmd {
	return r.cmd.ObjectEncoding(key)
}

func (r *Cmdable) ObjectIdleTime(key string) *redis.DurationCmd {
	return r.cmd.ObjectIdleTime(key)
}

func (r *Cmdable) Persist(key string) *redis.BoolCmd {
	return r.cmd.Persist(key)
}

func (r *Cmdable) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	return r.cmd.PExpire(key, expiration)
}

func (r *Cmdable) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return r.cmd.PExpireAt(key, tm)
}

func (r *Cmdable) PTTL(key string) *redis.DurationCmd {
	return r.cmd.PTTL(key)
}

func (r *Cmdable) RandomKey() *redis.StringCmd {
	return r.cmd.RandomKey()
}

func (r *Cmdable) Rename(key, newkey string) *redis.StatusCmd {
	return r.cmd.Rename(key, newkey)
}

func (r *Cmdable) RenameNX(key, newkey string) *redis.BoolCmd {
	return r.cmd.RenameNX(key, newkey)
}

func (r *Cmdable) Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	return r.cmd.Restore(key, ttl, value)
}

func (r *Cmdable) RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	return r.cmd.RestoreReplace(key, ttl, value)
}

func (r *Cmdable) Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	return r.cmd.Sort(key, sort)
}

func (r *Cmdable) SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	return r.cmd.SortStore(key, store, sort)
}

func (r *Cmdable) SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	return r.cmd.SortInterfaces(key, sort)
}
func (r *Cmdable) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.cmd.ZScan(key, cursor, match, count)
}

func (r *Cmdable) Touch(keys ...string) *redis.IntCmd {
	return r.cmd.Touch(keys...)
}

func (r *Cmdable) TTL(key string) *redis.DurationCmd {
	return r.cmd.TTL(key)
}

func (r *Cmdable) Type(key string) *redis.StatusCmd {
	return r.cmd.Type(key)
}

func (r *Cmdable) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.cmd.Scan(cursor, match, count)
}

func (r *Cmdable) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.cmd.SScan(key, cursor, match, count)
}

func (r *Cmdable) Append(key, value string) *redis.IntCmd {
	return r.cmd.Append(key, value)
}

func (r *Cmdable) BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	return r.cmd.BitCount(key, bitCount)
}

func (r *Cmdable) BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	return r.cmd.BitOpAnd(destKey, keys...)
}

func (r *Cmdable) BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	return r.cmd.BitOpOr(destKey, keys...)
}

func (r *Cmdable) BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	return r.cmd.BitOpXor(destKey, keys...)
}

func (r *Cmdable) BitOpNot(destKey string, key string) *redis.IntCmd {
	return r.cmd.BitOpNot(destKey, key)
}

func (r *Cmdable) BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	return r.cmd.BitPos(key, bit, pos...)
}

func (r *Cmdable) Decr(key string) *redis.IntCmd {
	return r.cmd.Decr(key)
}

func (r *Cmdable) DecrBy(key string, decrement int64) *redis.IntCmd {
	return r.cmd.DecrBy(key, decrement)
}

func (r *Cmdable) Get(key string) *redis.StringCmd {
	return r.cmd.Get(key)
}

func (r *Cmdable) GetBit(key string, offset int64) *redis.IntCmd {
	return r.cmd.GetBit(key, offset)
}

func (r *Cmdable) GetRange(key string, start, end int64) *redis.StringCmd {
	return r.cmd.GetRange(key, start, end)
}

func (r *Cmdable) GetSet(key string, value interface{}) *redis.StringCmd {
	return r.cmd.GetSet(key, value)
}

func (r *Cmdable) Incr(key string) *redis.IntCmd {
	return r.cmd.Incr(key)
}

func (r *Cmdable) IncrBy(key string, value int64) *redis.IntCmd {
	return r.cmd.IncrBy(key, value)
}

func (r *Cmdable) IncrByFloat(key string, value float64) *redis.FloatCmd {
	return r.cmd.IncrByFloat(key, value)
}

func (r *Cmdable) MGet(keys ...string) *redis.SliceCmd {
	return r.cmd.MGet(keys...)
}

func (r *Cmdable) MSet(pairs ...interface{}) *redis.StatusCmd {
	return r.cmd.MSet(pairs...)
}

func (r *Cmdable) MSetNX(pairs ...interface{}) *redis.BoolCmd {
	return r.cmd.MSetNX(pairs...)
}

func (r *Cmdable) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.cmd.Set(key, value, expiration)
}

func (r *Cmdable) SetBit(key string, offset int64, value int) *redis.IntCmd {
	return r.cmd.SetBit(key, offset, value)
}

func (r *Cmdable) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return r.cmd.SetNX(key, value, expiration)
}

func (r *Cmdable) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return r.cmd.SetXX(key, value, expiration)
}

func (r *Cmdable) SetRange(key string, offset int64, value string) *redis.IntCmd {
	return r.cmd.SetRange(key, offset, value)
}

func (r *Cmdable) StrLen(key string) *redis.IntCmd {
	return r.cmd.StrLen(key)
}

func (r *Cmdable) HDel(key string, fields ...string) *redis.IntCmd {
	return r.cmd.HDel(key, fields...)
}

func (r *Cmdable) HExists(key, field string) *redis.BoolCmd {
	return r.cmd.HExists(key, field)
}

func (r *Cmdable) HGet(key, field string) *redis.StringCmd {
	return r.cmd.HGet(key, field)
}

func (r *Cmdable) HGetAll(key string) *redis.StringStringMapCmd {
	return r.cmd.HGetAll(key)
}

func (r *Cmdable) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return r.cmd.HIncrBy(key, field, incr)
}

func (r *Cmdable) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return r.cmd.HIncrByFloat(key, field, incr)
}

func (r *Cmdable) HKeys(key string) *redis.StringSliceCmd {
	return r.cmd.HKeys(key)
}

func (r *Cmdable) HLen(key string) *redis.IntCmd {
	return r.cmd.HLen(key)
}

func (r *Cmdable) HMGet(key string, fields ...string) *redis.SliceCmd {
	return r.cmd.HMGet(key, fields...)
}

func (r *Cmdable) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	return r.cmd.HMSet(key, fields)
}

func (r *Cmdable) HSet(key, field string, value interface{}) *redis.BoolCmd {
	return r.cmd.HSet(key, field, value)
}

func (r *Cmdable) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	return r.cmd.HSetNX(key, field, value)
}

func (r *Cmdable) HVals(key string) *redis.StringSliceCmd {
	return r.cmd.HVals(key)
}

func (r *Cmdable) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.cmd.BLPop(timeout, keys...)
}

func (r *Cmdable) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.cmd.BRPop(timeout, keys...)
}

func (r *Cmdable) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	return r.cmd.BRPopLPush(source, destination, timeout)
}

func (r *Cmdable) LIndex(key string, index int64) *redis.StringCmd {
	return r.cmd.LIndex(key, index)
}

func (r *Cmdable) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	return r.cmd.LInsert(key, op, pivot, value)
}

func (r *Cmdable) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	return r.cmd.LInsertBefore(key, pivot, value)
}

func (r *Cmdable) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	return r.cmd.LInsertAfter(key, pivot, value)
}

func (r *Cmdable) LLen(key string) *redis.IntCmd {
	return r.cmd.LLen(key)
}

func (r *Cmdable) LPop(key string) *redis.StringCmd {
	return r.cmd.LPop(key)
}

func (r *Cmdable) LPush(key string, values ...interface{}) *redis.IntCmd {
	return r.cmd.LPush(key, values...)
}

func (r *Cmdable) LPushX(key string, value interface{}) *redis.IntCmd {
	return r.cmd.LPushX(key, value)
}

func (r *Cmdable) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.cmd.LRange(key, start, stop)
}

func (r *Cmdable) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	return r.cmd.LRem(key, count, value)
}

func (r *Cmdable) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	return r.cmd.LSet(key, index, value)
}

func (r *Cmdable) LTrim(key string, start, stop int64) *redis.StatusCmd {
	return r.cmd.LTrim(key, start, stop)
}

func (r *Cmdable) RPop(key string) *redis.StringCmd {
	return r.cmd.RPop(key)
}

func (r *Cmdable) RPopLPush(source, destination string) *redis.StringCmd {
	return r.cmd.RPopLPush(source, destination)
}

func (r *Cmdable) RPush(key string, values ...interface{}) *redis.IntCmd {
	return r.cmd.RPush(key, values...)
}

func (r *Cmdable) RPushX(key string, value interface{}) *redis.IntCmd {
	return r.cmd.RPushX(key, value)
}

func (r *Cmdable) SAdd(key string, members ...interface{}) *redis.IntCmd {
	return r.cmd.SAdd(key, members...)
}

func (r *Cmdable) SCard(key string) *redis.IntCmd {
	return r.cmd.SCard(key)
}

func (r *Cmdable) SDiff(keys ...string) *redis.StringSliceCmd {
	return r.cmd.SDiff(keys...)
}

func (r *Cmdable) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	return r.cmd.SDiffStore(destination, keys...)
}

func (r *Cmdable) SInter(keys ...string) *redis.StringSliceCmd {
	return r.cmd.SInter(keys...)
}

func (r *Cmdable) SInterStore(destination string, keys ...string) *redis.IntCmd {
	return r.cmd.SInterStore(destination, keys...)
}

func (r *Cmdable) SIsMember(key string, member interface{}) *redis.BoolCmd {
	return r.cmd.SIsMember(key, member)
}

func (r *Cmdable) SMembers(key string) *redis.StringSliceCmd {
	return r.cmd.SMembers(key)
}

func (r *Cmdable) SMembersMap(key string) *redis.StringStructMapCmd {
	return r.cmd.SMembersMap(key)
}

func (r *Cmdable) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	return r.cmd.SMove(source, destination, member)
}

func (r *Cmdable) SPop(key string) *redis.StringCmd {
	return r.cmd.SPop(key)
}

func (r *Cmdable) SPopN(key string, count int64) *redis.StringSliceCmd {
	return r.cmd.SPopN(key, count)
}

func (r *Cmdable) SRandMember(key string) *redis.StringCmd {
	return r.cmd.SRandMember(key)
}

func (r *Cmdable) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	return r.cmd.SRandMemberN(key, count)
}

func (r *Cmdable) SRem(key string, members ...interface{}) *redis.IntCmd {
	return r.cmd.SRem(key, members...)
}

func (r *Cmdable) SUnion(keys ...string) *redis.StringSliceCmd {
	return r.cmd.SUnion(keys...)
}

func (r *Cmdable) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	return r.cmd.SUnionStore(destination, keys...)
}

func (r *Cmdable) XAdd(a *redis.XAddArgs) *redis.StringCmd {
	return r.cmd.XAdd(a)
}

func (r *Cmdable) XDel(stream string, ids ...string) *redis.IntCmd {
	return r.cmd.XDel(stream, ids...)
}

func (r *Cmdable) XLen(stream string) *redis.IntCmd {
	return r.cmd.XLen(stream)
}

func (r *Cmdable) XRange(stream, start, stop string) *redis.XMessageSliceCmd {
	return r.cmd.XRange(stream, start, stop)
}

func (r *Cmdable) XRangeN(stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	return r.cmd.XRangeN(stream, start, stop, count)
}

func (r *Cmdable) XRevRange(stream string, start, stop string) *redis.XMessageSliceCmd {
	return r.cmd.XRevRange(stream, start, stop)
}

func (r *Cmdable) XRevRangeN(stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	return r.cmd.XRevRangeN(stream, start, stop, count)
}

func (r *Cmdable) XRead(a *redis.XReadArgs) *redis.XStreamSliceCmd {
	return r.cmd.XRead(a)
}

func (r *Cmdable) XReadStreams(streams ...string) *redis.XStreamSliceCmd {
	return r.cmd.XReadStreams(streams...)
}

func (r *Cmdable) XGroupCreate(stream, group, start string) *redis.StatusCmd {
	return r.cmd.XGroupCreate(stream, group, start)
}

func (r *Cmdable) XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd {
	return r.cmd.XGroupCreateMkStream(stream, group, start)
}

func (r *Cmdable) XGroupSetID(stream, group, start string) *redis.StatusCmd {
	return r.cmd.XGroupSetID(stream, group, start)
}

func (r *Cmdable) XGroupDestroy(stream, group string) *redis.IntCmd {
	return r.cmd.XGroupDestroy(stream, group)
}

func (r *Cmdable) XGroupDelConsumer(stream, group, consumer string) *redis.IntCmd {
	return r.cmd.XGroupDelConsumer(stream, group, consumer)
}

func (r *Cmdable) XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	return r.cmd.XReadGroup(a)
}

func (r *Cmdable) XAck(stream, group string, ids ...string) *redis.IntCmd {
	return r.cmd.XAck(stream, group, ids...)
}

func (r *Cmdable) XPending(stream, group string) *redis.XPendingCmd {
	return r.cmd.XPending(stream, group)
}

func (r *Cmdable) XPendingExt(a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	return r.cmd.XPendingExt(a)
}

func (r *Cmdable) XClaim(a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	return r.cmd.XClaim(a)
}

func (r *Cmdable) XClaimJustID(a *redis.XClaimArgs) *redis.StringSliceCmd {
	return r.cmd.XClaimJustID(a)
}

func (r *Cmdable) XTrim(key string, maxLen int64) *redis.IntCmd {
	return r.cmd.XTrim(key, maxLen)
}

func (r *Cmdable) XTrimApprox(key string, maxLen int64) *redis.IntCmd {
	return r.cmd.XTrimApprox(key, maxLen)
}

func (r *Cmdable) BZPopMax(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return r.cmd.BZPopMax(timeout, keys...)
}

func (r *Cmdable) BZPopMin(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return r.cmd.BZPopMin(timeout, keys...)
}

func (r *Cmdable) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return r.cmd.ZAdd(key, members...)
}

func (r *Cmdable) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	return r.cmd.ZAddNX(key, members...)
}

func (r *Cmdable) ZAddXX(key string, members ...redis.Z) *redis.IntCmd {
	return r.cmd.ZAddXX(key, members...)
}

func (r *Cmdable) ZAddCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.cmd.ZAddCh(key, members...)
}

func (r *Cmdable) ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.cmd.ZAddNXCh(key, members...)
}

func (r *Cmdable) ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.cmd.ZAddXXCh(key, members...)
}

func (r *Cmdable) ZIncr(key string, member redis.Z) *redis.FloatCmd {
	return r.cmd.ZIncr(key, member)
}

func (r *Cmdable) ZIncrNX(key string, member redis.Z) *redis.FloatCmd {
	return r.cmd.ZIncrNX(key, member)
}

func (r *Cmdable) ZIncrXX(key string, member redis.Z) *redis.FloatCmd {
	return r.cmd.ZIncrXX(key, member)
}

func (r *Cmdable) ZCard(key string) *redis.IntCmd {
	return r.cmd.ZCard(key)
}

func (r *Cmdable) ZCount(key, min, max string) *redis.IntCmd {
	return r.cmd.ZCount(key, min, max)
}

func (r *Cmdable) ZLexCount(key, min, max string) *redis.IntCmd {
	return r.cmd.ZLexCount(key, min, max)
}

func (r *Cmdable) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	return r.cmd.ZIncrBy(key, increment, member)
}

func (r *Cmdable) ZInterStore(destination string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return r.cmd.ZInterStore(destination, store, keys...)
}

func (r *Cmdable) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	return r.cmd.ZPopMax(key, count...)
}

func (r *Cmdable) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	return r.cmd.ZPopMin(key, count...)
}

func (r *Cmdable) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.cmd.ZRange(key, start, stop)
}

func (r *Cmdable) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return r.cmd.ZRangeWithScores(key, start, stop)
}

func (r *Cmdable) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.cmd.ZRangeByScore(key, opt)
}

func (r *Cmdable) ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.cmd.ZRangeByLex(key, opt)
}

func (r *Cmdable) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return r.cmd.ZRangeByScoreWithScores(key, opt)
}

func (r *Cmdable) ZRank(key, member string) *redis.IntCmd {
	return r.cmd.ZRank(key, member)
}

func (r *Cmdable) ZRem(key string, members ...interface{}) *redis.IntCmd {
	return r.cmd.ZRem(key, members...)
}

func (r *Cmdable) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	return r.cmd.ZRemRangeByRank(key, start, stop)
}

func (r *Cmdable) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	return r.cmd.ZRemRangeByScore(key, min, max)
}

func (r *Cmdable) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	return r.cmd.ZRemRangeByLex(key, min, max)
}

func (r *Cmdable) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.cmd.ZRevRange(key, start, stop)
}

func (r *Cmdable) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return r.cmd.ZRevRangeWithScores(key, start, stop)
}

func (r *Cmdable) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.cmd.ZRevRangeByScore(key, opt)
}

func (r *Cmdable) ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.cmd.ZRevRangeByLex(key, opt)
}

func (r *Cmdable) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return r.cmd.ZRevRangeByScoreWithScores(key, opt)
}

func (r *Cmdable) ZRevRank(key, member string) *redis.IntCmd {
	return r.cmd.ZRevRank(key, member)
}

func (r *Cmdable) ZScore(key, member string) *redis.FloatCmd {
	return r.cmd.ZScore(key, member)
}

func (r *Cmdable) ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return r.cmd.ZUnionStore(dest, store, keys...)
}

func (r *Cmdable) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	return r.cmd.PFAdd(key, els...)
}

func (r *Cmdable) PFCount(keys ...string) *redis.IntCmd {
	return r.cmd.PFCount(keys...)
}

func (r *Cmdable) PFMerge(dest string, keys ...string) *redis.StatusCmd {
	return r.cmd.PFMerge(dest, keys...)
}

func (r *Cmdable) BgRewriteAOF() *redis.StatusCmd {
	return r.cmd.BgRewriteAOF()
}

func (r *Cmdable) BgSave() *redis.StatusCmd {
	return r.cmd.BgSave()
}

func (r *Cmdable) ClientKill(ipPort string) *redis.StatusCmd {
	return r.cmd.ClientKill(ipPort)
}

func (r *Cmdable) ClientKillByFilter(keys ...string) *redis.IntCmd {
	return r.cmd.ClientKillByFilter(keys...)
}

func (r *Cmdable) ClientList() *redis.StringCmd {
	return r.cmd.ClientList()
}

func (r *Cmdable) ClientPause(dur time.Duration) *redis.BoolCmd {
	return r.cmd.ClientPause(dur)
}

func (r *Cmdable) ClientID() *redis.IntCmd {
	return r.cmd.ClientID()
}

func (r *Cmdable) ConfigGet(parameter string) *redis.SliceCmd {
	return r.cmd.ConfigGet(parameter)
}

func (r *Cmdable) ConfigResetStat() *redis.StatusCmd {
	return r.cmd.ConfigResetStat()
}

func (r *Cmdable) ConfigSet(parameter, value string) *redis.StatusCmd {
	return r.cmd.ConfigSet(parameter, value)
}

func (r *Cmdable) ConfigRewrite() *redis.StatusCmd {
	return r.cmd.ConfigRewrite()
}

func (r *Cmdable) DBSize() *redis.IntCmd {
	return r.cmd.DBSize()
}

func (r *Cmdable) FlushAll() *redis.StatusCmd {
	return r.cmd.FlushAll()
}

func (r *Cmdable) FlushAllAsync() *redis.StatusCmd {
	return r.cmd.FlushAllAsync()
}

func (r *Cmdable) FlushDB() *redis.StatusCmd {
	return r.cmd.FlushDB()
}

func (r *Cmdable) FlushDBAsync() *redis.StatusCmd {
	return r.cmd.FlushDBAsync()
}

func (r *Cmdable) Info(section ...string) *redis.StringCmd {
	return r.cmd.Info(section...)
}

func (r *Cmdable) LastSave() *redis.IntCmd {
	return r.cmd.LastSave()
}

func (r *Cmdable) Save() *redis.StatusCmd {
	return r.cmd.Save()
}

func (r *Cmdable) Shutdown() *redis.StatusCmd {
	return r.cmd.Shutdown()
}

func (r *Cmdable) ShutdownSave() *redis.StatusCmd {
	return r.cmd.ShutdownSave()
}

func (r *Cmdable) ShutdownNoSave() *redis.StatusCmd {
	return r.cmd.ShutdownNoSave()
}

func (r *Cmdable) SlaveOf(host, port string) *redis.StatusCmd {
	return r.cmd.SlaveOf(host, port)
}

func (r *Cmdable) Time() *redis.TimeCmd {
	return r.cmd.Time()
}

func (r *Cmdable) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	return r.cmd.Eval(script, keys, args...)
}

func (r *Cmdable) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return r.cmd.EvalSha(sha1, keys, args...)
}

func (r *Cmdable) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	return r.cmd.ScriptExists(hashes...)
}

func (r *Cmdable) ScriptFlush() *redis.StatusCmd {
	return r.cmd.ScriptFlush()
}

func (r *Cmdable) ScriptKill() *redis.StatusCmd {
	return r.cmd.ScriptKill()
}

func (r *Cmdable) ScriptLoad(script string) *redis.StringCmd {
	return r.cmd.ScriptLoad(script)
}

func (r *Cmdable) DebugObject(key string) *redis.StringCmd {
	return r.cmd.DebugObject(key)
}

func (r *Cmdable) Publish(channel string, message interface{}) *redis.IntCmd {
	return r.cmd.Publish(channel, message)
}

func (r *Cmdable) PubSubChannels(pattern string) *redis.StringSliceCmd {
	return r.cmd.PubSubChannels(pattern)
}

func (r *Cmdable) PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	return r.cmd.PubSubNumSub(channels...)
}

func (r *Cmdable) PubSubNumPat() *redis.IntCmd {
	return r.cmd.PubSubNumPat()
}

func (r *Cmdable) ClusterSlots() *redis.ClusterSlotsCmd {
	return r.cmd.ClusterSlots()
}

func (r *Cmdable) ClusterNodes() *redis.StringCmd {
	return r.cmd.ClusterNodes()
}

func (r *Cmdable) ClusterMeet(host, port string) *redis.StatusCmd {
	return r.cmd.ClusterMeet(host, port)
}

func (r *Cmdable) ClusterForget(nodeID string) *redis.StatusCmd {
	return r.cmd.ClusterForget(nodeID)
}

func (r *Cmdable) ClusterReplicate(nodeID string) *redis.StatusCmd {
	return r.cmd.ClusterReplicate(nodeID)
}

func (r *Cmdable) ClusterResetSoft() *redis.StatusCmd {
	return r.cmd.ClusterResetSoft()
}

func (r *Cmdable) ClusterResetHard() *redis.StatusCmd {
	return r.cmd.ClusterResetHard()
}

func (r *Cmdable) ClusterInfo() *redis.StringCmd {
	return r.cmd.ClusterInfo()
}

func (r *Cmdable) ClusterKeySlot(key string) *redis.IntCmd {
	return r.cmd.ClusterKeySlot(key)
}

func (r *Cmdable) ClusterGetKeysInSlot(slot int, count int) *redis.StringSliceCmd {
	return r.cmd.ClusterGetKeysInSlot(slot, count)
}

func (r *Cmdable) ClusterCountFailureReports(nodeID string) *redis.IntCmd {
	return r.cmd.ClusterCountFailureReports(nodeID)
}

func (r *Cmdable) ClusterCountKeysInSlot(slot int) *redis.IntCmd {
	return r.cmd.ClusterCountKeysInSlot(slot)
}

func (r *Cmdable) ClusterDelSlots(slots ...int) *redis.StatusCmd {
	return r.cmd.ClusterDelSlots(slots...)
}

func (r *Cmdable) ClusterDelSlotsRange(min, max int) *redis.StatusCmd {
	return r.cmd.ClusterDelSlotsRange(min, max)
}

func (r *Cmdable) ClusterSaveConfig() *redis.StatusCmd {
	return r.cmd.ClusterSaveConfig()
}

func (r *Cmdable) ClusterSlaves(nodeID string) *redis.StringSliceCmd {
	return r.cmd.ClusterSlaves(nodeID)
}

func (r *Cmdable) ClusterFailover() *redis.StatusCmd {
	return r.cmd.ClusterFailover()
}

func (r *Cmdable) ClusterAddSlots(slots ...int) *redis.StatusCmd {
	return r.cmd.ClusterAddSlots(slots...)
}

func (r *Cmdable) ClusterAddSlotsRange(min, max int) *redis.StatusCmd {
	return r.cmd.ClusterAddSlotsRange(min, max)
}

func (r *Cmdable) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return r.cmd.GeoAdd(key, geoLocation...)
}

func (r *Cmdable) GeoPos(key string, members ...string) *redis.GeoPosCmd {
	return r.cmd.GeoPos(key, members...)
}

func (r *Cmdable) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.cmd.GeoRadius(key, longitude, latitude, query)
}

func (r *Cmdable) GeoRadiusRO(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.cmd.GeoRadiusRO(key, longitude, latitude, query)
}

func (r *Cmdable) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.cmd.GeoRadiusByMember(key, member, query)
}

func (r *Cmdable) GeoRadiusByMemberRO(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.cmd.GeoRadiusByMemberRO(key, member, query)
}

func (r *Cmdable) GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	return r.cmd.GeoDist(key, member1, member2, unit)
}

func (r *Cmdable) GeoHash(key string, members ...string) *redis.StringSliceCmd {
	return r.cmd.GeoHash(key, members...)
}

func (r *Cmdable) ReadOnly() *redis.StatusCmd {
	return r.cmd.ReadOnly()
}

func (r *Cmdable) ReadWrite() *redis.StatusCmd {
	return r.cmd.ReadWrite()
}

func (r *Cmdable) MemoryUsage(key string, samples ...int) *redis.IntCmd {
	return r.cmd.MemoryUsage(key, samples...)
}
