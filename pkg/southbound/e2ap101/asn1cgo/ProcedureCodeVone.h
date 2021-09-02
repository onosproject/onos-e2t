/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-CommonDataTypes"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_ProcedureCodeVone_H_
#define	_ProcedureCodeVone_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"

#ifdef __cplusplus
extern "C" {
#endif

/* ProcedureCodeVone */
typedef long	 ProcedureCodeVone_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_ProcedureCodeVone_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_ProcedureCodeVone;
asn_struct_free_f ProcedureCodeVone_free;
asn_struct_print_f ProcedureCodeVone_print;
asn_constr_check_f ProcedureCodeVone_constraint;
ber_type_decoder_f ProcedureCodeVone_decode_ber;
der_type_encoder_f ProcedureCodeVone_encode_der;
xer_type_decoder_f ProcedureCodeVone_decode_xer;
xer_type_encoder_f ProcedureCodeVone_encode_xer;
per_type_decoder_f ProcedureCodeVone_decode_uper;
per_type_encoder_f ProcedureCodeVone_encode_uper;
per_type_decoder_f ProcedureCodeVone_decode_aper;
per_type_encoder_f ProcedureCodeVone_encode_aper;
#define ProcedureCodeVone_id_E2setupVone	((ProcedureCodeVone_t)1)
#define ProcedureCodeVone_id_ErrorIndicationVone	((ProcedureCodeVone_t)2)
#define ProcedureCodeVone_id_ResetVone	((ProcedureCodeVone_t)3)
#define ProcedureCodeVone_id_RICcontrolVone	((ProcedureCodeVone_t)4)
#define ProcedureCodeVone_id_RICindicationVone	((ProcedureCodeVone_t)5)
#define ProcedureCodeVone_id_RICserviceQueryVone	((ProcedureCodeVone_t)6)
#define ProcedureCodeVone_id_RICserviceUpdateVone	((ProcedureCodeVone_t)7)
#define ProcedureCodeVone_id_RICsubscriptionVone	((ProcedureCodeVone_t)8)
#define ProcedureCodeVone_id_RICsubscriptionDeleteVone	((ProcedureCodeVone_t)9)
#define ProcedureCodeVone_id_E2nodeConfigurationUpdateVone	((ProcedureCodeVone_t)10)
#define ProcedureCodeVone_id_E2connectionUpdateVone	((ProcedureCodeVone_t)11)

#ifdef __cplusplus
}
#endif

#endif	/* _ProcedureCodeVone_H_ */
#include "asn_internal.h"
