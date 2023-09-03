## NS: Nameservice Module Demo

## Concepts
Nameservice resolves a human readable addresses to an account address. 

The nameservice module enables users to bid on a name in a perpetual auction. If it has not previously been purchased, then the first bidder becomes the owner of the name. 
If there is a previous owner, then the user must bid a higher price than the current owner paid.

## State

Collections
- Names
- Owner

Index
- Names by Owner
- Names by Resolve Address

## Messages

**Services**
```protobuf
service Msg {
  rpc Reserve(MsgReserve) returns (MsgReserveResponse);
}

service Query {

  rpc Name(QueryNameRequest) returns (QueryNameResponse) {
    option (google.api.http).get = "/cosmos/ns/name/{name}";
  }
}
```

**Messages**
```protobuf
message MsgReserve {
  option (cosmos.msg.v1.signer) = "owner";

  // name defines the human readable address
  string name = 1;

  // resolveAddress defines the bech32 address to resolve to
  string resolveAddress = 2 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];

  // owner is the address of the owner of listing
  string owner = 3 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];

  //   price is the last price paid for listing
  repeated cosmos.base.v1beta1.Coin amount = 4 [
    (gogoproto.nullable) = false,
    (amino.dont_omitempty) = true,
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"
  ];
}

message QueryNameRequest { 
  string name = 1; 
}
```

## Client

### Tx
`appd tx nameservice [name] [resolve_address] [amount] [flags]`

### Query
`appd q nameservice whois [name]`
