package win

//ref PVOID
//ref DWORD
//ref PCERT_INFO
//ref HCERTSTORE
//ref PCCERT_CONTEXT

type PFN_CRYPT_GET_SIGNER_CERTIFICATE func(pvGetArg PVOID, dwCertEncodingType DWORD, pSignerId PCERT_INFO, hMsgCertStore HCERTSTORE) PCCERT_CONTEXT