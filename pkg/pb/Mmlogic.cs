// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: api/mmlogic.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace OpenMatch {

  /// <summary>Holder for reflection information generated from api/mmlogic.proto</summary>
  public static partial class MmlogicReflection {

    #region Descriptor
    /// <summary>File descriptor for api/mmlogic.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MmlogicReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChFhcGkvbW1sb2dpYy5wcm90bxIDYXBpGhJhcGkvbWVzc2FnZXMucHJvdG8a",
            "HGdvb2dsZS9hcGkvYW5ub3RhdGlvbnMucHJvdG8aLHByb3RvYy1nZW4tc3dh",
            "Z2dlci9vcHRpb25zL2Fubm90YXRpb25zLnByb3RvIi4KE1F1ZXJ5VGlja2V0",
            "c1JlcXVlc3QSFwoEcG9vbBgBIAEoCzIJLmFwaS5Qb29sIjQKFFF1ZXJ5VGlj",
            "a2V0c1Jlc3BvbnNlEhwKB3RpY2tldHMYASADKAsyCy5hcGkuVGlja2V0MnYK",
            "B01tTG9naWMSawoMUXVlcnlUaWNrZXRzEhguYXBpLlF1ZXJ5VGlja2V0c1Jl",
            "cXVlc3QaGS5hcGkuUXVlcnlUaWNrZXRzUmVzcG9uc2UiJILT5JMCHiIZL3Yx",
            "L21tbG9naWMvdGlja2V0czpxdWVyeToBKjABQpgDWiBvcGVuLW1hdGNoLmRl",
            "di9vcGVuLW1hdGNoL3BrZy9wYqoCCU9wZW5NYXRjaJJB5gISvwEKFU1NIExv",
            "Z2ljIChEYXRhIExheWVyKSJJCgpPcGVuIE1hdGNoEhZodHRwczovL29wZW4t",
            "bWF0Y2guZGV2GiNvcGVuLW1hdGNoLWRpc2N1c3NAZ29vZ2xlZ3JvdXBzLmNv",
            "bSpWChJBcGFjaGUgMi4wIExpY2Vuc2USQGh0dHBzOi8vZ2l0aHViLmNvbS9n",
            "b29nbGVmb3JnYW1lcy9vcGVuLW1hdGNoL2Jsb2IvbWFzdGVyL0xJQ0VOU0Uy",
            "AzEuMCoCAQIyEGFwcGxpY2F0aW9uL2pzb246EGFwcGxpY2F0aW9uL2pzb25S",
            "OwoDNDA0EjQKKlJldHVybmVkIHdoZW4gdGhlIHJlc291cmNlIGRvZXMgbm90",
            "IGV4aXN0LhIGCgSaAgEHcj0KGE9wZW4gTWF0Y2ggRG9jdW1lbnRhdGlvbhIh",
            "aHR0cHM6Ly9vcGVuLW1hdGNoLmRldi9zaXRlL2RvY3MvYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::OpenMatch.MessagesReflection.Descriptor, global::Google.Api.AnnotationsReflection.Descriptor, global::Grpc.Gateway.ProtocGenSwagger.Options.AnnotationsReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::OpenMatch.QueryTicketsRequest), global::OpenMatch.QueryTicketsRequest.Parser, new[]{ "Pool" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::OpenMatch.QueryTicketsResponse), global::OpenMatch.QueryTicketsResponse.Parser, new[]{ "Tickets" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class QueryTicketsRequest : pb::IMessage<QueryTicketsRequest> {
    private static readonly pb::MessageParser<QueryTicketsRequest> _parser = new pb::MessageParser<QueryTicketsRequest>(() => new QueryTicketsRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<QueryTicketsRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::OpenMatch.MmlogicReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public QueryTicketsRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public QueryTicketsRequest(QueryTicketsRequest other) : this() {
      pool_ = other.pool_ != null ? other.pool_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public QueryTicketsRequest Clone() {
      return new QueryTicketsRequest(this);
    }

    /// <summary>Field number for the "pool" field.</summary>
    public const int PoolFieldNumber = 1;
    private global::OpenMatch.Pool pool_;
    /// <summary>
    /// The Pool representing the set of Filters to be queried.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::OpenMatch.Pool Pool {
      get { return pool_; }
      set {
        pool_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as QueryTicketsRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(QueryTicketsRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Pool, other.Pool)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (pool_ != null) hash ^= Pool.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (pool_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Pool);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (pool_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Pool);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(QueryTicketsRequest other) {
      if (other == null) {
        return;
      }
      if (other.pool_ != null) {
        if (pool_ == null) {
          Pool = new global::OpenMatch.Pool();
        }
        Pool.MergeFrom(other.Pool);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (pool_ == null) {
              Pool = new global::OpenMatch.Pool();
            }
            input.ReadMessage(Pool);
            break;
          }
        }
      }
    }

  }

  public sealed partial class QueryTicketsResponse : pb::IMessage<QueryTicketsResponse> {
    private static readonly pb::MessageParser<QueryTicketsResponse> _parser = new pb::MessageParser<QueryTicketsResponse>(() => new QueryTicketsResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<QueryTicketsResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::OpenMatch.MmlogicReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public QueryTicketsResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public QueryTicketsResponse(QueryTicketsResponse other) : this() {
      tickets_ = other.tickets_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public QueryTicketsResponse Clone() {
      return new QueryTicketsResponse(this);
    }

    /// <summary>Field number for the "tickets" field.</summary>
    public const int TicketsFieldNumber = 1;
    private static readonly pb::FieldCodec<global::OpenMatch.Ticket> _repeated_tickets_codec
        = pb::FieldCodec.ForMessage(10, global::OpenMatch.Ticket.Parser);
    private readonly pbc::RepeatedField<global::OpenMatch.Ticket> tickets_ = new pbc::RepeatedField<global::OpenMatch.Ticket>();
    /// <summary>
    /// The Tickets that meet the Filter criteria requested by the Pool.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::OpenMatch.Ticket> Tickets {
      get { return tickets_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as QueryTicketsResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(QueryTicketsResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!tickets_.Equals(other.tickets_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= tickets_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      tickets_.WriteTo(output, _repeated_tickets_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += tickets_.CalculateSize(_repeated_tickets_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(QueryTicketsResponse other) {
      if (other == null) {
        return;
      }
      tickets_.Add(other.tickets_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            tickets_.AddEntriesFrom(input, _repeated_tickets_codec);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
