GRPC SERVER & FUSE SERVER:
    - Both need to have mutual locking
    - util/dist_locks


Currnetly:
    - Grpc Server using util/dist_locks
    - Hooks using util/dist_locks
    - Separate looks on Fuse

GO FAIL Locks vs BLOCK Locks
