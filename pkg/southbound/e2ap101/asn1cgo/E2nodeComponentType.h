/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2nodeComponentType_H_
#define	_E2nodeComponentType_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum E2nodeComponentType {
	E2nodeComponentType_gNB	= 0,
	E2nodeComponentType_gNB_CU_UP	= 1,
	E2nodeComponentType_gNB_DU	= 2,
	E2nodeComponentType_en_gNB	= 3,
	E2nodeComponentType_eNB	= 4,
	E2nodeComponentType_ng_eNB	= 5
	/*
	 * Enumeration is extensible
	 */
} e_E2nodeComponentType;

/* E2nodeComponentType */
typedef long	 E2nodeComponentType_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_E2nodeComponentType_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_E2nodeComponentType;
extern const asn_INTEGER_specifics_t asn_SPC_E2nodeComponentType_specs_1;
asn_struct_free_f E2nodeComponentType_free;
asn_struct_print_f E2nodeComponentType_print;
asn_constr_check_f E2nodeComponentType_constraint;
ber_type_decoder_f E2nodeComponentType_decode_ber;
der_type_encoder_f E2nodeComponentType_encode_der;
xer_type_decoder_f E2nodeComponentType_decode_xer;
xer_type_encoder_f E2nodeComponentType_encode_xer;
per_type_decoder_f E2nodeComponentType_decode_uper;
per_type_encoder_f E2nodeComponentType_encode_uper;
per_type_decoder_f E2nodeComponentType_decode_aper;
per_type_encoder_f E2nodeComponentType_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _E2nodeComponentType_H_ */
#include "asn_internal.h"
